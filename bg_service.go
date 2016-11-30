package main

import "runtime"

const (
	RUNNING = "RUNNING"
	IDLE	= "IDLE"

)

type (
	Thread struct {
		ServiceName	string
		EventTrigger	[]string
		Function	func()
		Status		string
	}

	ThreadPool struct {
		Threads		[]*Thread
	}
)

var ThreadPool ThreadPool

func (j *Thread) New(serviceName string, eventTrigger []string) *Thread{
	j.ServiceName = serviceName
	j.EventTrigger = eventTrigger
	j.Status = IDLE
	return j
}

func (j *Thread) implementFunction(){
	os := runtime.GOOS
	for {
		select{
		case :
		}
	}
}

func (j *Thread) AddToThreadPool(){
	ThreadPool.Threads = append(ThreadPool.Threads, j)
}

func (j *Thread) Run(){
	go j.Function()
	j.Status = RUNNING
}

