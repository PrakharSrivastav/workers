package main

import (
	"fmt"
	"github.com/PrakharSrivastav/workers/dispatcher"
	"github.com/PrakharSrivastav/workers/worker"
	"testing"
	"time"
)

var terms = []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}

func BenchmarkConcurrent(b *testing.B) {
	dd := dispatcher.New().Start(8) // start up worker pool
	for n := 0; n < b.N; n++ {
		for i := range terms {
			dd.Submit(worker.Job{
				ID:        i,
				Name:      fmt.Sprintf("JobID::%d", i),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	dd := dispatcher.New().Start(8) // start up worker pool
	for n := 0; n < b.N; n++ {
		for i := range terms {
			dd.Submit(worker.Job{
				ID:        i,
				Name:      fmt.Sprintf("JobID::%d", i),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}
}
