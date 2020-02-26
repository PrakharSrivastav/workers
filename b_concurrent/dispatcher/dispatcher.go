package dispatcher

import (
	"github.com/PrakharSrivastav/workers/b_concurrent/worker"
)

// New returns a new dispatcher. A Dispatcher communicates between the client
// and the worker. Its main job is to receive a job and share it on the WorkPool
// WorkPool is the link between the dispatcher and all the workers as
// the WorkPool of the dispatcher is common JobPool for all the workers
func New() *disp {
	return &disp{
		WorkChan: make(worker.JobChannel),
		Queue: make(worker.JobQueue),
	}
}

// disp is the link between the client and the workers
type disp struct {
	Workers  []*worker.Worker     // this is the list of workers that dispatcher tracks
	WorkChan worker.JobChannel      // client submits job to this channel
	Queue worker.JobQueue // this is the shared JobPool between the workers
}

// Start creates pool of num count of workers.
func (d *disp) Start(num int) *disp {
	d.Workers = make([]*worker.Worker, num)
	for i := 1; i <= num; i++ {
		wrk := worker.New(i, make(worker.JobChannel), d.Queue, make(chan struct{}))
		wrk.Start()
		d.Workers = append(d.Workers, wrk)
	}
	go d.process()
	return d
}

// process listens to a job submitted on WorkChan and
// relays it to the WorkPool. The WorkPool is shared between
// the workers.
func (d *disp) process() {
	for {
		select {
		case job := <-d.WorkChan: // listen to any submitted job on the WorkChan
			// wait for a worker to submit JobChan to Queue
			// note that this Queue is shared among all workers.
			// Whenever there is an available JobChan on Queue pull it
			jobChan := <-d.Queue

			// Once a jobChan is available, send the submitted Job on this JobChan
			jobChan <- job
		}
	}
}

func (d *disp) Submit(job worker.Job) {
	d.WorkChan <- job
}
