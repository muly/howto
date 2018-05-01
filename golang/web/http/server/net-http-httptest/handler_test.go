package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler(t *testing.T) {

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.URL)
	//.Println(req.URL)

	resp := httptest.NewRecorder()
	handler(resp, req)

	t.Log(resp.Body)
	t.Log(resp.Code)

	if resp.Code != http.StatusOK {
		t.Errorf("response test failed. wanted %v, but got %v", http.StatusOK, resp.Code)
		return
	}

	//TODO: need to add reasponse body test
}
