package main

import (
	"fmt"
	"os"
)

const(
	FILENAME = "settings.json"
)

var setting Settings

func main(){
	setting.ReadSettings(FILENAME)
	var pc_pool []ProcessGroup
	for _, procName := range setting.Process{
		pc := ProcessGroup{}
		pc.Name = procName
		pc.Init()
		pc.OnDie = func(){
			fmt.Println("Chrome.exe was killed")
			os.Exit(0)
		}
		pc_pool = append(pc_pool, pc)
	}

	for _, procg := range pc_pool{
		procg.Watch()
	}

	for{

	}
}