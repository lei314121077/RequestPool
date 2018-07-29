package main

import (
	"poolone"
	"fmt"
	"strconv"
	"time"
)

var (
	//MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue  = os.Getenv("MAX_QUEUE")
	MaxWorker = 100000 // this max work lenght is 100000
	MaxQueue  = 100000  // this jobQueue chan max lenght is 100000
)


//During our web server initialization we create a Dispatcher and call Run()
// to create the pool of workers and to start listening for jobs that would appear in the JobQueue.
func initdata(){
	//创建一个Dispatcher并调用Run()以创建工作池并开始侦听将出现在其中的作业JobQueue。
	dispatcher := poolone.NewDispatcher(MaxWorker)
	//　init this jobqueue chan lenght is MaxQueue
	poolone.JobQueue = make(chan poolone.Job, MaxQueue)
	dispatcher.Run()
}

func main(){
	start := time.Now()
	fmt.Printf(start.Format("2006-01-02 15:04:05"),"\n")
	initdata()
	for i:=0; i< MaxWorker; i++{
		p := poolone.Payload{
			Name:fmt.Sprintf("工作者-[%s]",strconv.Itoa(i)),
		}
		poolone.JobQueue <- poolone.Job{
			Payload:p,
		}
		//time.Sleep(time.Second)
	}
	close(poolone.JobQueue)
	//记录结束时间
	end := time.Since(start)
	fmt.Printf("共计耗时:%v", end)
}
