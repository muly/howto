// demonstrate a simple app with docker (with small image size)

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

//// multi-step container build:
// compile go code and add the binary into the container
//		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
//		docker build -t hellosmallapp .
// docker run to create the container from the specified image (using tag in below example)
// 		docker run --rm hellosmallapp

// https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/ #Part 3: Compile!
// vs
// https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2 #Way 2: Build Outside your Dockerfile
