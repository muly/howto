package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_helloHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello", helloHandler)
	// router.Handle("GET", "/hello", helloHandler)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status: wanted %v, but got %v", http.StatusOK, w.Code)
	}
}

// references:
// see this blog for more complex examples: https://medium.com/@gauravsingharoy/build-your-first-api-server-with-httprouter-in-golang-732b7b01f6ab
