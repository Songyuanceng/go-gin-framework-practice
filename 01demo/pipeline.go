package main

import (
	"log"
	"sync"
)

func HasClosed(c <-chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

type SyncFlag interface {
	Wait()
	Chan() <-chan struct{}
	Done() bool
}

func NewSyncFlag() (done func(), flag SyncFlag) {
	f := &syncFlag{
		c: make(chan struct{}),
	}
	return f.done, f
}

type syncFlag struct {
	once sync.Once
	c    chan struct{}
}

func (f *syncFlag) done() {
	f.once.Do(func() {
		close(f.c)
	})
}

func (f *syncFlag) Wait() {
	<-f.c
}

func (f *syncFlag) Chan() <-chan struct{} {
	return f.c
}

func (f *syncFlag) Done() bool {
	return HasClosed(f.c)
}

type pipelineThread struct {
	sigs         []chan struct{}
	chanExit     chan struct{}
	interrupt    SyncFlag
	setInterrupt func()
	err          error
}

func newPipelineThread(l int) *pipelineThread {
	p := &pipelineThread{
		sigs:     make([]chan struct{}, l),
		chanExit: make(chan struct{}),
	}
	p.setInterrupt, p.interrupt = NewSyncFlag()

	for i := range p.sigs {
		p.sigs[i] = make(chan struct{})
	}
	return p
}

type Pipeline struct {
	mtx         sync.Mutex
	workerChans []chan struct{}
	prevThd     *pipelineThread
}

//创建流水线，参数个数是每个任务的子过程数，每个参数对应子过程的并发度。
func NewPipeline(workers ...int) *Pipeline {
	if len(workers) < 1 {
		panic("NewPipeline need aleast one argument")
	}

	workersChan := make([]chan struct{}, len(workers))
	for i := range workersChan {
		workersChan[i] = make(chan struct{}, workers[i])
	}

	prevThd := newPipelineThread(len(workers))
	for _, sig := range prevThd.sigs {
		close(sig)
	}
	close(prevThd.chanExit)

	return &Pipeline{
		workerChans: workersChan,
		prevThd:     prevThd,
	}
}

//往流水线推入一个任务。如果第一个步骤的并发数达到设定上限，这个函数会堵塞等待。
//如果流水线中有其它任务失败（返回非nil），任务不被执行，函数返回false。
func (p *Pipeline) Async(works ...func() error) bool {
	if len(works) != len(p.workerChans) {
		panic("Async: arguments number not matched to NewPipeline(...)")
	}

	p.mtx.Lock()
	if p.prevThd.interrupt.Done() {
		p.mtx.Unlock()
		return false
	}
	prevThd := p.prevThd
	thisThd := newPipelineThread(len(p.workerChans))
	p.prevThd = thisThd
	p.mtx.Unlock()

	lock := func(idx int) bool {
		select {
		case <-prevThd.interrupt.Chan():
			return false
		case <-prevThd.sigs[idx]: //wait for signal
		}
		select {
		case <-prevThd.interrupt.Chan():
			return false
		case p.workerChans[idx] <- struct{}{}: //get lock
		}
		return true
	}
	if !lock(0) {
		thisThd.setInterrupt()
		<-prevThd.chanExit
		thisThd.err = prevThd.err
		close(thisThd.chanExit)
		return false
	}
	go func() { //watch interrupt of previous thread
		select {
		case <-prevThd.interrupt.Chan():
			thisThd.setInterrupt()
		case <-thisThd.chanExit:
		}
	}()
	go func() {
		var err error
		for i, work := range works {
			close(thisThd.sigs[i]) //signal next thread
			if work != nil {
				err = work()
			}
			if err != nil || (i+1 < len(works) && !lock(i+1)) {
				thisThd.setInterrupt()
				break
			}
			<-p.workerChans[i] //release lock
		}

		<-prevThd.chanExit
		if prevThd.interrupt.Done() {
			thisThd.setInterrupt()
		}
		if prevThd.err != nil {
			thisThd.err = prevThd.err
		} else {
			thisThd.err = err
		}
		close(thisThd.chanExit)
	}()
	return true
}

//等待流水线中所有任务执行完毕或失败，返回第一个错误，如果无错误则返回nil。
func (p *Pipeline) Wait() error {
	p.mtx.Lock()
	lastThd := p.prevThd
	p.mtx.Unlock()
	<-lastThd.chanExit
	return lastThd.err
}

func main() {
	//恢复上次执行的checkpoint，如果是第一次执行就获取一个初始值。
	checkpoint := loadCheckpoint()

	//工序(1)在pipeline外执行，最后一个工序是保存checkpoint
	pipeline := NewPipeline(4, 8, 2, 1)
	for {
		//(1)
		//加载100条数据，并修改变量checkpoint
		//data是数组，每个元素是一条评论，之后的联表、NLP都直接修改data里的每条记录。
		data, err := extractReviewsFromA(&checkpoint, 100)
		if err != nil {
			log.Print(err)
			break
		}

		//这里有个Golang著名的坑。
		//“checkpoint”是循环体外的变量，它在内存中只有一个实例并在循环中不断被修改，所以不能在异步中使用它。
		//这里创建一个副本curCheckpoint，储存本次循环的checkpoint。
		curCheckpoint := checkpoint

		ok := pipeline.Async(func() error {
			//(2)
			return joinUserFromB(data)
		}, func() error {
			//(3)
			return nlp(data)
		}, func() error {
			//(4)
			return loadDataToC(data)
		}, func() error {
			//(5)保存checkpoint
			log.Print("done:", curCheckpoint)
			return saveCheckpoint(curCheckpoint)
		})
		if !ok {
			break
		}

		if len(data) < 100 {
			break
		} //处理完毕
	}
	err := pipeline.Wait()
	if err != nil {
		log.Print(err)
	}
}
