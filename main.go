package main

import "fmt"

const(
	FILENAME = "settings.json"
)

func main(){
	var setting Settings
	setting.ReadSettings(FILENAME)
	var pc_pool []ProcessCatcher
	for _, procName := range setting.Process{
		pc := ProcessCatcher{}
		pc.Name = procName
		pc.Init()
		pc_pool = append(pc_pool, pc)
	}

	for _, proc := range pc_pool[0].Processes{
		fmt.Println(proc.PID)
	}
}