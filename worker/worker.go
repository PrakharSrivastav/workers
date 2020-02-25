package worker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Job represents a single entity that should be processed.
// For example a struct that should be saved to database
type Job struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Worker is a a single processor. Typically its possible to
// start multiple workers for better throughput
type Worker struct {
	ID      int           // id of the worker
	JobChan chan Job      // a channel to receive single unit of work
	JobPool chan chan Job // pool of job channels. Used for distributing work between workers
	Quit    chan struct{} // a channel to quit working
}

func New(ID int, JobChan chan Job, JobPool chan chan Job, Quit chan struct{}) *Worker{
	return &Worker{
		ID:      ID,
		JobChan: JobChan,
		JobPool: JobPool,
		Quit:    Quit,
	}
}

func (wr *Worker) Start() {
	c := &http.Client{Timeout: time.Millisecond * 1000}
	go func() {
		for {
			// Just put the job in the job pool so that any
			// select that is unblocked can process
			wr.JobPool <- wr.JobChan
			select {
			case job := <-wr.JobChan:
				callApi(job.ID,wr.ID,c)
			case <-wr.Quit:
				return
			}
		}
	}()
}

// stop closes the Quit channel on the worker.
func (wr *Worker) Stop() {
	close(wr.Quit)
}


func callApi(num,id int , c *http.Client) {
	baseURL := "https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d"

	ur := fmt.Sprintf(baseURL, num)
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		fmt.Printf("error creating a request for term %s :: error is %+v", num, err)
		return
	}
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("error querying for term %s :: error is %+v", num, err)
		return
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body :: error is %+v", num, err)
		return
	}
	log.Printf("%d  :: ok", id)
}