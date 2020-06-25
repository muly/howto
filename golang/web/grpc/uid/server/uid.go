package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/globalsign/mgo"
)

type uid struct {
	db        *mgo.Collection
	session   *mgo.Session
	mainStart int64
	increment int
	sync.Mutex
}

const wait = 10 * time.Millisecond // need to add this to config

var state uid

func genUniqueIDs(num int64) []string {
	ids := make([]string, 0, num)

	for i := num; i > 0; i-- {
		now := time.Now().Unix()
		ids = append(ids, getUniqueID(now))
	}

	return ids
}

func getUniqueID(now int64) string {
	state.Lock()
	defer state.Unlock()
	// once the 12 bits are filled, wait for next second and reset the counter
	if state.increment >= 4095 { // Note: 4095 is the max num corresponding to the 12 bits allocated for the sequence number for uid to be used within the second.
		waitUntillNextSec(state.mainStart)
		now = time.Now().Unix()
		state.increment = 0
	}

	// once the next second is reached, reset the counter
	if state.mainStart != now {
		state.mainStart = now
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
		time.Sleep(wait)
	}
}

func dbInsert(key string) error {
	//TODO: need to reduce this to single function call instead of 2 (find & upsert)
	//insert should fail if the key already exists
	if n, err := state.db.FindId(key).Count(); err != nil {
		return err
	} else if n > 0 {
		return fmt.Errorf("key %s already exists in DB", key)
	}

	// insert records
	_, err := state.db.UpsertId(key, struct{}{})
	if err != nil {
		return err
	}

	return nil
}
