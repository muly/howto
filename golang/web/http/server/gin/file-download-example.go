// input url example: localhost:8080/dl
// output required is a file named "download-file.csv" downloaded with the content "hello,world"

package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.GET("/dl", downloadFile)

	router.Run(":8080")
}

func downloadFile(c *gin.Context) {
	data := bytes.NewBuffer([]byte("hello,world"))

	c.Header("Content-Disposition", "attachment; filename=download-file.csv")
	c.Data(http.StatusOK, "text/csv", data.Bytes())
}
