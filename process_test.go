package main

import (
	"testing"
	"strings"
	"fmt"
)

func TestWinTaskListCalling(t *testing.T){
	p := ProcessCatcher{}
	p.Name = "chrome.exe"
	out := p.winTaskList()

	fmt.Println(out)

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
	p.Name = "chrome.exe"
	out := p.winTaskList()
	processElem, _ := p.readProcessList(out)
	fmt.Println(processElem)
}
