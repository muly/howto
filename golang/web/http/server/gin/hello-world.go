// simple hello world router example using gin framework
// input url example-TODO: localhost:8080/hello
// output in html response body expected is: hello world

//Resources:
// https://phalt.co/a-simple-api-in-go/

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func showHelloPage(c *gin.Context) {

	content := gin.H{"Hello": "world"}

	c.JSON(200, content)
}

func main() {

	router = gin.Default()

	router.GET("/hello", showHelloPage)

	// Start serving the application
	router.Run(":8080")
}
