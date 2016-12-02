//[ProcessWatcherGroup]
//	==>	[ProcessWatcher]	==>	go intervalChecker()	->	WatcherDeleteChannel
//	==>	[ProcessWatcher]	==>	go intervalChecker()	->
//	==>	WatcherDeleteChannel	-> find & delete()

package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"net/http"
	"log"
)

const(
	WINDOWS = "windows"
)

type(
	ProcessWatcherGroup        struct {
		Name                             string
		Processes                        []Watcher
		WatcherDeleteChannel             chan int
		ProcessWatcherGroupDeleteChannel chan bool
		NewProcessListenerStatus	bool
		DeleteWatcherHandlerStatus	bool
		Setting				ProcessSetting
	}

	Watcher struct{
		ProcessName	string
		PID		int
		parent		*ProcessWatcherGroup
	}
)

func (t *Watcher) read(in string){
	pattern := `([0-9]+)`
	exp := regexp.MustCompile(pattern)
	items := exp.FindAllString(in, -1)
	PID, err := strconv.Atoi(items[0])
	GoPanic(err)
	t.PID = PID
}

func  intervalChecking(pid int, procg *ProcessWatcherGroup){
	defer fmt.Println("PID ", pid, " Killed")
	for{
		fmt.Println("i am alive", pid, procg.Name)
		isRunning := findProcess(pid, procg.Name)
		if !isRunning {
			procg.WatcherDeleteChannel <- pid
			return
		}
		time.Sleep(time.Duration(setting.IntervalTime) * time.Millisecond)
	}
}

func (t *ProcessWatcherGroup) ProcessAlreadyInCollector(pid int) bool{
	for _, proc := range t.Processes{
		if proc.PID == pid{
			return true
		}
	}
	return false
}

func NewProcessListener(t *ProcessWatcherGroup){
	defer fmt.Println(t.Name, "NewProcessListener Killed")
	defer func(){
		t.NewProcessListenerStatus = false
	}()

	for{

		process := t.readProcessList(t.winTaskList())

		if len(t.Processes) < len(process){
			for _, proc := range process{
				if !t.ProcessAlreadyInCollector(proc.PID){
					t.Processes = append(t.Processes, proc)
					go intervalChecking(proc.PID, t)
					break
				}
			}

			if t.DeleteWatcherHandlerStatus == false{
				t.DeleteWatcherHandlerStatus = true
				go DeleteWatcherHandler(t)
			}
			fmt.Println("new process created")
		}

		time.Sleep(time.Duration(setting.IntervalTime) * time.Millisecond)
	}
}

func DeleteWatcherHandler(t *ProcessWatcherGroup){
	defer fmt.Println(t.Name, " DeleteProcessHandler Killed")
	defer func(){
		t.DeleteWatcherHandlerStatus = false
	}()

	for {
		if len(t.Processes) == 0 {
			t.Die()
			return
		}
		select {
		case PID := <- t.WatcherDeleteChannel:
			t.deleteFromPool(PID)
			if len(t.Processes) == 0{
				t.Die()
				return
			}
		}
		time.Sleep(time.Duration(setting.IntervalTime) * time.Millisecond)
	}
}

func (t *ProcessWatcherGroup) Watch(){

	for _, process := range t.Processes{
		go intervalChecking(process.PID, t)
	}

	//ProcessWatcherGroup thread looking for new process created
	t.NewProcessListenerStatus = true
	go NewProcessListener(t)

	//ProcessWatcherGroup thread waiting for someone killed then kill the his watcher
	t.DeleteWatcherHandlerStatus = true
	go DeleteWatcherHandler(t)
}

func (t *ProcessWatcherGroup) Die() {
	for _, event := range t.Setting.Event{
		url := fmt.Sprintf(BASEURL, event, setting.Key, t.Name)
		_, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		//data, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("call to", url)
	}
}

func (t *ProcessWatcherGroup) Init(){
	if setting.OS == WINDOWS {
		out := t.winTaskList()
		elems := t.readProcessList(out)
		for _, item := range elems {
			item.parent = t
			t.Processes = append(t.Processes, item)
		}
	}

	t.WatcherDeleteChannel = make(chan int)
}

func (t *ProcessWatcherGroup) deleteFromPool(pid int) bool{
	var i int
	for i = 0; i < len(t.Processes); i++{
		item := t.Processes[i]
		if item.PID == pid {
			t.Processes = append(t.Processes[:i], t.Processes[i+1:]...)
			return true
		}
	}
	return false
}

func (t *ProcessWatcherGroup) winTaskList() string {
	exeParamProcess := fmt.Sprintf("imagename eq %s", t.Name)
	command := exec.Command("tasklist", "-fi", exeParamProcess)
	out, err := command.Output()
	GoPanic(err)
	return string(out)
}

func (t *ProcessWatcherGroup) readProcessList(out string) []Watcher {
	pattern := fmt.Sprintf(`%s([ ]+)([0-9]+)([ ]+)([A-Za-z]+)([ ]+)([0-9])([ ]+) ([0-9,]+) K`, t.Name)
	exp := regexp.MustCompile(pattern)
	items := exp.FindAllStringSubmatch(out, -1)
	var processes []Watcher
	for _, item := range items{
		var process Watcher
		process.read(item[0])
		processes = append(processes, process)
	}
	return processes
}

func findProcess(pid int, processName string) bool{
	if setting.OS == WINDOWS{
		procg := ProcessWatcherGroup{Name:processName}
		procs := procg.readProcessList(procg.winTaskList())
		for _, proc := range procs{
			if proc.PID == pid {
				return true
			}
		}
	}
	return false
}




