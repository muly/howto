package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Resp struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}


func Handler() (Resp, error) {
	output := Resp{ID: 1, Date: time.Now().String()}
	fmt.Println(output)
	return output, nil
}

func main() {
	lambda.Start(Handler)
}

//  GOOS=linux go build -o lamda
