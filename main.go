package main

import (
	"fmt"
	"runtime"
	"time"
	"net/http"
	"log"
	"io/ioutil"
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
	for _, procName := range setting.Process{
		pc := ProcessWatcherGroup{}
		pc.Name = procName.Name
		pc.Init()
		pc.OnDie = func(){
			fmt.Println(pc.Name, " was fully killed")
			for _, event := range procName.Event{
				fmt.Println("call event ", event)
				url := fmt.Sprintf(BASEURL, event, setting.Key, pc.Name)
				resp, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
					return
				}
				data, _ := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()

				fmt.Println(string(data))
			}

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