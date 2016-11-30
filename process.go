package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"strconv"
	"runtime"
	"os"
	"time"
)

type(
	ProcessGroup        struct {
		Name		string
		OnDie		func()
		Processes	[]Process
	}

	Process struct{
		ProcessName	string
		PID		int
		parent		*ProcessGroup
	}
)

func (t *Process) read(in string){
	pattern := `([0-9]+)`
	exp := regexp.MustCompile(pattern)
	items := exp.FindAllString(in, -1)
	PID, err := strconv.Atoi(items[0])
	GoPanic(err)
	t.PID = PID
}

func (t * Process) intervalChecking(){
	defer t.parent.deleteFromPool(t.PID)
	for{
		_, err := os.FindProcess(t.PID)
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Duration(setting.IntervalTime) * time.Millisecond)
	}
}

func (t *ProcessGroup) Watch(){
	for _, process := range t.Processes{
		go process.intervalChecking()
	}

	go func(){
		for len(t.Processes) != 0{
			fmt.Println(len(t.Processes))
		}
		t.OnDie()
	}()
}

func (t *ProcessGroup) Init(){
	if runtime.GOOS == "windows" {
		out := t.winTaskList()
		elems, err := t.readProcessList(out)
		GoPanic(err)
		for _, item := range elems {
			item.parent = t
			t.Processes = append(t.Processes, item)
		}
	}
}

func (t * ProcessGroup) deleteFromPool(pid int) bool{
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

func (t *ProcessGroup) winTaskList() string {
	exeParamProcess := fmt.Sprintf("imagename eq %s", t.Name)
	command := exec.Command("tasklist", "-fi", exeParamProcess)
	out, err := command.Output()
	GoPanic(err)
	return string(out)
}

func (t *ProcessGroup) readProcessList(out string) ([]Process, error){
	pattern := fmt.Sprintf(`%s([ ]+)([0-9]+)([ ]+)([A-Za-z]+)([ ]+)([0-9])([ ]+) ([0-9,]+) K`, t.Name)
	exp := regexp.MustCompile(pattern)
	items := exp.FindAllStringSubmatch(out, -1)
	var processes []Process
	for _, item := range items{
		var process Process
		process.read(item[0])
		processes = append(processes, process)
	}
	return processes, nil
}




