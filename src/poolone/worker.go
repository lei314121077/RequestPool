package poolone

//var (
//	MaxWorker = os.Getenv("MAX_WORKERS")
//	MaxQueue  = os.Getenv("MAX_QUEUE")
//)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	Name string
	WorkerPool  chan chan Job
	JobChannel  chan Job
	quit    	chan bool
}

func NewWorker(workerPool chan chan Job, name string) Worker {
	//fmt.Printf("create a worker, name is :%s \n",name)

	return Worker{
		Name:name,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			//fmt.Printf("[%s]get a work ,process task~~![%s] \n",w.Name, len(w.JobChannel))

			select {

			case job := <- w.JobChannel:
				// we have received a work request. need process work task!
				//fmt.Printf("[%s]->we have received a work request. need process work task! task lenght! [%d] !\n",w.Name, len(w.WorkerPool))

				job.Payload.Process()

				//time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)


			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

