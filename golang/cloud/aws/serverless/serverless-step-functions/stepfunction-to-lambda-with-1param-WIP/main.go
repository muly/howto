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

type Req struct{
	A int `json:"A"`
}

func Handler(req  Req) (Resp, error) {
	output := Resp{ID: req.A, Date: time.Now().String()}
	fmt.Println(output)
	return output, nil
}

func main() {
	lambda.Start(Handler)
}

// npm install --save-dev serverless-step-functions 
// GOOS=linux go build -o lambdaExecutable
// sls deploy