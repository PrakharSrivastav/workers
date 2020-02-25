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
		WorkChan: make(chan worker.Job),
		WorkPool: make(chan chan worker.Job),
	}
}

// disp is the link between the client and the workers
type disp struct {
	Workers  []*worker.Worker     // this is the list of workers that dispatcher tracks
	WorkChan chan worker.Job      // client submits job to this channel
	WorkPool chan chan worker.Job // this is the shared JobPool between the workers
}

// Start creates pool of num count of workers.
func (d *disp) Start(num int) *disp {
	d.Workers = make([]*worker.Worker, num)
	for i := 1; i <= num; i++ {
		wrk := worker.New(i, make(chan worker.Job), d.WorkPool, make(chan struct{}))
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
		case j := <-d.WorkChan: // listen to a submitted job on WorkChannel
			work := <-d.WorkPool // pull out a chan chan jobs
			work <- j            // share the chan job on the pool
		}
	}
}

func (d *disp) Submit(job worker.Job) {
	d.WorkChan <- job
}

