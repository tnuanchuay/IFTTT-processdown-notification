package main

import (
	"fmt"
	"runtime"
	"time"
)

const(
	FILENAME = "settings.json"
)

var setting Settings

func main(){
	setting.ReadSettings(FILENAME)
	setting.OS = runtime.GOOS
	var pc_pool []ProcessWatcherGroup
	for _, procName := range setting.Process{
		pc := ProcessWatcherGroup{}
		pc.Name = procName
		pc.Init()
		pc.OnDie = func(){
			fmt.Printf("%s was fully killed\n", pc.Name)
			//os.Exit(0)
		}
		pc_pool = append(pc_pool, pc)
	}

	for _, procg := range pc_pool{
		procg.Watch()
	}

	for{
		time.Sleep(10 * time.Second)
	}
}