package main

import (
	"os/exec"
	"fmt"
	"regexp"
)

type(
	ProcessCatcher	struct {
		Name		string
		PID		int
		OnDie		func()
	}

	Process struct{
		ProcessName	string
		PID		int
	}
)

func (t *ProcessCatcher) winTaskList() string {
	exeParamProcess := fmt.Sprintf("imagename eq %s", t.Name)
	command := exec.Command("tasklist", "-fi", exeParamProcess)
	out, err := command.Output()
	GoPanic(err)
	return string(out)
}

func (t *ProcessCatcher) readProcessList(out string) ([][]string, error){
	pattern := fmt.Sprintf("%s([ ]+)([0-9]+)([ ]+)([A-Za-z])([ ]+)([0-9])([ ]+)([0-9,]+) K", t.Name)
	exp := regexp.MustCompile(pattern)
	elements := exp.FindAllStringSubmatch(out, -1)

	return elements, nil
}




