// example to demonstrate how to connect to an aws sqs


package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {

	access := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("REGION") // "us-east-1"

	token := ""
	creds := credentials.NewStaticCredentials(access, secret, token)

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess, err := session.NewSession(cfg)
	if err != nil {
		fmt.Println("session failed:", err)
	}

	// Create a SQS service client.
	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success")
	// As these are pointers, printing them out directly would not be useful.
	for i, urls := range result.QueueUrls {
		// Avoid dereferencing a nil pointer.
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}

}
