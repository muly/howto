package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func Test_Router(t *testing.T){
	h := mux.NewRouter()
	h.HandleFunc("/", helloworld)
	go http.ListenAndServe(":8080", h)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil{
		t.Errorf("failed to get:%v", err)
		return
	}
	if resp.StatusCode != http.StatusOK{
		t.Errorf("expected status code 200 but received %v", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		t.Logf("failed to read body:%v", err)
		return
	}
	t.Logf("received body %s", string(b))
}