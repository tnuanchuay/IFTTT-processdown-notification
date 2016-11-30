package main

import "os"

func GoPanic(err error){
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
