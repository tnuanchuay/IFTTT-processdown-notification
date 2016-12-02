package main

import (
	"runtime"
	"time"
)

const(
	FILENAME	=	"settings.json"
	BASEURL		=	"https://maker.ifttt.com/trigger/%s/with/key/%s?value1=%s"
)

var setting Settings

func main(){
	setting.ReadSettings(FILENAME)
	setting.OS = runtime.GOOS
	var pc_pool []ProcessWatcherGroup
	for _, procSetting := range setting.Process{
		pc := ProcessWatcherGroup{Name:procSetting.Name, Setting:procSetting}
		pc.Init()
		pc_pool = append(pc_pool, pc)
	}

	for i := 0; i < len(pc_pool); i++{
		pc_pool[i].Watch()
	}

	for{
		time.Sleep(10 * time.Second)
	}
}