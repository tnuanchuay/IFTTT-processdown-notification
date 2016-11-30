package main


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

var threadsPool ThreadPool

func (j *Thread) New(serviceName string, eventTrigger []string) *Thread{
	j.ServiceName = serviceName
	j.EventTrigger = eventTrigger
	j.Status = IDLE
	return j
}

func (j *Thread) implementFunction(){

}

func (j *Thread) AddToThreadPool(){
	threadsPool.Threads = append(threadsPool.Threads, j)
}

func (j *Thread) Run(){
	go j.Function()
	j.Status = RUNNING
}

