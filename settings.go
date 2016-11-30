package main

import (
	"io/ioutil"
	"encoding/json"
)

type(
	Settings struct {
		Process		[]string	`json:"processes"`
		Key		string		`json:"ifttt-maker-key"`
		Event		[]string	`json:"event"`
	}
)


func (j *Settings) ReadSettings(filename string){
	byteFile, err := ioutil.ReadFile(filename)
	GoPanic(err)

	err = json.Unmarshal(byteFile, j)
	GoPanic(err)
}
