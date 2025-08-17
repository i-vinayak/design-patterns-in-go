package main

import (
	"fmt"
	"sync"
)

type data struct {
	message string
}

var singletonData *data
var once sync.Once

// 1. Implementation using sync
func NewSingleton(msg string) *data {
	once.Do(func() {
		fmt.Println("Creating instance")
		singletonData = &data{message: msg}
	})
	return singletonData
}

// 2. Implementation using init
// func init() {
// 	if singletonData == nil {
// 		singletonData = &data{}
// 	}
// }

// func NewSingleton(msg string) *data {
// 	if singletonData != nil {
// 		fmt.Println("Singleton Already exists")
// 	}
// 	return singletonData
// }

// var lock sync.Mutex

// 3. Implementation using Lock
// func NewSingleton(msg string) *data {
// 	if singletonData == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singletonData == nil {
// 			singletonData = &data{message: msg}
// 		} else {
// 			fmt.Println("Already exists")
// 		}
// 	} else {
// 		fmt.Println("Instance exists")
// 	}
// 	return singletonData
// }
