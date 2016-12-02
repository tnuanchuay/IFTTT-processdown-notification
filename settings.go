package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type(
	Settings struct {
		Process		[]ProcessSetting	`json:"processes-watcher"`
		Key		string			`json:"ifttt-maker-key"`
		IntervalTime	int			`json:"interval-time"`
		OS		string
	}

	ProcessSetting struct{
		Name		string			`json:"name"`
		Event		[]string		`json:"event"`
	}
)


func (j *Settings) ReadSettings(filename string){
	byteFile, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Do you copy settings.json.env to settings.json?, if you don't, it's in src file")
		GoPanic(err)
	}

	err = json.Unmarshal(byteFile, j)
	GoPanic(err)
}
