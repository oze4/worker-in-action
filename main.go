package main

import (
    "github.com/oze4/worker"
)

func main() {
    jobs := []worker.Job{
        AddedByUser{name: "1"},
        AddedByUser{name: "2"},
        AddedByUser{name: "3"},
        AddedByUser{name: "4"},
        AddedByUser{name: "5"},
        AddedByUser{name: "6"},
    }
    
    worker.Do(jobs, 5)
}

type AddedByUser struct {
	name string
}

func (a AddedByUser) Name() string {
	return a.name
}

func (a AddedByUser) Callback() worker.JobResponse {
	// User added func/callback goes here
	return worker.JobResponse{}
}