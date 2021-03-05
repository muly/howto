package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Req struct {
	OutputLocation string `json:"OutputLocation"`
}

func Handler(input Req) error {
	fmt.Println("OutputLocation: ", input.OutputLocation)

	return nil
}

func main() {

	lambda.Start(Handler)

}

/* //// DEPLOYMENT STEPS ////
0. build for linux and zip the executable
	- GOOS=linux go build -o lamda lambda-1-param.go
*/
