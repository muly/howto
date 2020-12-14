// demonstrate a simple app with docker (with small image size)

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("hello world", uuid.New())
}

//// multi-step container build:
// pre-steps: make sure all the dependencies are downloaded, as the script in this Dockerfile expects the dependecies to be in place.
// 		go mod init
// 		go mod vendor
// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
//		docker build -t hellosmallapp .
// docker run to create the container from the specified image (using tag in below example)
// 		docker run --rm hellosmallapp

// references:
// https://fabianlee.org/2020/01/26/golang-using-multi-stage-builds-to-create-clean-docker-images/
//
//
// https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/ #Part 3: Compile!
// vs
// https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2 #Way 2: Build Outside your Dockerfile
