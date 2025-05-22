package main

import "sync"

type data struct {
	message string
}

var singletonData *data
var once sync.Once

func NewSingleton(msg string) *data {
	once.Do(func() {
		singletonData = &data{message: msg}
	})
	return singletonData
}
