// demonstrate to do the below:
// register an api gateway (for example as a trigger to lambda function)
// choose the security as "AWS IAM"
// create a client in go and call the api gateway endpoint
//
// note: to trigger this endpoint from Postman, choose "AWS Signature" for Authorization. ref https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-use-postman-to-call-api.html

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	sign "github.com/AdRoll/goamz/aws" // used to sign the request with AWS IAM security
)

func main() {

	// below are the credentials of the user that has access to API Gateway related policies like "AmazonAPIGatewayInvokeFullAccess"
	access := os.Getenv("AWS_API_GATEWAY_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_API_GATEWAY_SECRET_ACCESS_KEY")
	//is the URL generated after the API Gateway endpoint is created (for example as a trigger step for Lambda)
	reqURL := os.Getenv("AWS_API_GATEWAY_URL")
	region := os.Getenv("REGION")
	serviceName := "execute-api" // service name is required as "execute-api" when making the api gateway service call with AWS IAM security

	auth := sign.Auth{AccessKey: access, SecretKey: secret}
	v4sign := sign.NewV4Signer(auth, serviceName, sign.Region{Name: region})

	r, err := http.NewRequest("GET", reqURL, nil) // Choose the http method and body per the need
	if err != nil {
		fmt.Println(err)
		return
	}

	v4sign.Sign(r)

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
	fmt.Println("api gateway endpoint executed successfully:", string(respBody))

}

/*
// when the api key env variables are not having the keys info. you may see error like below
	unsupported protocol scheme ""
//
*/
