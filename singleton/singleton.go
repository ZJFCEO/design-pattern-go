package main

import "sync"

type Single struct {
	data int
}

var singleton *Single
var once sync.Once

func GetInstance() *Single {
	once.Do(func() {
		singleton = &Single{data: 100}
	})
	//singleton = &Single{data: 100}
	return singleton
}
