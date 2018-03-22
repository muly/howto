// demonstrate a simple app with docker

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
// 		docker build -t helloapp .
// docker run to create the container from the specified image (using tag in below example)
// 		docker run --rm helloapp
