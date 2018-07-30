package poolone

import (
	"fmt"
	"strconv"
		)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	Name string 			 	// 调度的名字
	MaxWorkers int 			 	// 获取 调试的大小
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		Name:"Diskpatcher Mangage:",
		MaxWorkers:maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {

		worker := NewWorker(d.WorkerPool, fmt.Sprintf("this is %s work",strconv.Itoa(i)))
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			//　dispatch get a job work
			//fmt.Println("dispatch get a job work")
			//time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				if jobChannel, ok := <- d.WorkerPool; ok{
					// dispatch the job to the worker job channel
					jobChannel <- job
				}
			}(job)
		default:
			//fmt.Println("process ok!")
		}


	}
}

