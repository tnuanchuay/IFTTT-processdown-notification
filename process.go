package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"strconv"
	"runtime"
)

type(
	ProcessCatcher	struct {
		Name		string
		OnDie		func()
		Processes	[]Process
	}

	Process struct{
		ProcessName	string
		PID		int
	}
)

func (t * ProcessCatcher) Init(){
	if runtime.GOOS == "windows" {
		out := t.winTaskList()
		elems, err := t.readProcessList(out)
		GoPanic(err)
		for _, item := range elems {
			t.Processes = append(t.Processes, item)
		}
	}
}

func (t *Process) read(in string){
	pattern := `([0-9]+)`
	exp := regexp.MustCompile(pattern)
	items := exp.FindAllString(in, -1)
	PID, err := strconv.Atoi(items[0])
	GoPanic(err)
	t.PID = PID
}

func (t *ProcessCatcher) winTaskList() string {
	exeParamProcess := fmt.Sprintf("imagename eq %s", t.Name)
	command := exec.Command("tasklist", "-fi", exeParamProcess)
	out, err := command.Output()
	GoPanic(err)
	return string(out)
}

func (t *ProcessCatcher) readProcessList(out string) ([]Process, error){
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




