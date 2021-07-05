package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	
	"github.com/google/uuid"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
	tc "github.com/testcontainers/testcontainers-go"
)

var port string

func TestEndpointDockerfile(t *testing.T) {
	// // 1. setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker test: %s", err)
	}

	resource, err := pool.BuildAndRunWithOptions("./Dockerfile", &dockertest.RunOptions{
		Name: "my-api-image",
		PortBindings: map[dc.Port][]dc.PortBinding{
			"8080": {{HostPort: "8080"}},
		},
	})
	if err != nil {
		log.Fatalf("Could not build/run container: %s", err)
	}

	port = resource.GetPort("8080/tcp")

	// // 2. run tests
	RunTest(t)

	// // 3. cleanup
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestEndpointDockerCompose(t *testing.T) {
	// // 1. setup
	composeFilePaths := []string {"./docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := tc.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string {
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, err)
		return 
	}

	port = "8081"

	// // 2. run tests
	RunTest(t)

	// // 3. cleanup
	execError = compose.Down()
	err = execError.Error
	if err != nil {
		 fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, err)
		 return
	}
	return 
}


func RunTest(t *testing.T){
	for _, compress := range []bool{true, false} {
		if compress {
			t.Log("test case: compress")
		} else {
			t.Log("test case: no compress")
		}
		reqURL := fmt.Sprintf("http://localhost:%s/", port)
		r, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			t.Errorf("new request failed: %v", err)
			return
		}
		fmt.Println(r)

		if compress {
			r.Header.Set("Accept-Encoding", "gzip")
		}

		client := http.Client{}
		resp, err := client.Do(r)
		if err != nil {
			t.Errorf("get failed: %v", err)
			return
		}
		defer resp.Body.Close()
		if compress {
			if resp.Header.Get("Content-Encoding") != "gzip" {
				t.Errorf("Content-Encoding missing: %v", resp.Header.Get("Content-Encoding"))
			}
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("read body failed: %v", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("incorrect status code. expected %v, but got %v", http.StatusOK, resp.StatusCode)
			t.Errorf("received data: %s", string(data))
			return
		}

		t.Logf("header Content-Encoding: %s", resp.Header.Get("Content-Encoding"))
		t.Logf("received data len: %v", len(data))
	}
}