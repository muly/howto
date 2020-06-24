package main

import (
	"fmt"
	"sync"
	"time"
)

type uid struct {
	mainStart int64
	increment int
	sync.Mutex
}

var state uid

func genUniqueIDs(start time.Time, num int32) []string {
	var ids []string

	for i := int32(0); i < num; i++ {
		ids = append(ids, getUniqueID())
	}

	return ids
}

func getUniqueID() string {
	state.Lock()
	defer state.Unlock()
	// once the 12 bits are filled, wait for next second and reset the counter
	if state.increment == 4095 {
		waitUntillNextSec(state.mainStart)
		state.increment = 0
	}

	now := time.Now().Unix()

	// once the next second is reached, reset the counter
	if state.mainStart != now {
		state.mainStart = time.Now().Unix()
		state.increment = 0
	}

	state.increment++

	return fmt.Sprintf("%05x%x%03x", 0, state.mainStart, state.increment)
}

func waitUntillNextSec(mainStart int64) {
	for {
		now := time.Now().Unix()
		if mainStart != now {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func dbInsert(key string) error {
	//insert should fail if the key already exists
	//TODO: insert into mongodb

	return nil
}
