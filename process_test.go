package main

import (
	"testing"
	"strings"
	"os/exec"
	"os"
)

func TestWinTaskListCalling(t *testing.T){
	p := ProcessCatcher{}
	serviceName := "chrome.exe"
	exec.Command(serviceName).Run()
	p.Name = serviceName
	out := p.winTaskList()
	headTableWords := []string{"Image Name", "PID", "Session Name", "Session#", "Mem Usage"}

	for _, word := range headTableWords{
		isContain := strings.Contains(out, word)
		if isContain != true{
			t.Errorf(`It should be "found string %s" when call winSearch`, word)
		}
	}

}

func TestWinTaskListReading(t *testing.T){
	p := ProcessCatcher{}
	serviceName := "chrome.exe"
	exec.Command(serviceName).Run()
	p.Name = serviceName
	out := p.winTaskList()
	elms, _ := p.readProcessList(out)
	_, err := os.FindProcess(elms[0].PID)
	if err != nil{
		t.Error(err)
	}
}
