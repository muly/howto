// TODO: check for syntax errors

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var access = os.Getenv("AWS_ACCESS_KEY_ID")
var secret = os.Getenv("AWS_SECRET_ACCESS_KEY")
var region = os.Getenv("REGION")

func getAwsSession() (*session.Session, error) {

	token := ""
	creds := credentials.NewStaticCredentials(access, secret, token)

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess, err := session.NewSession(cfg)
	if err != nil {
		fmt.Println("session failed:", err)
		return nil, err
	}
	return sess, nil
}

func main() {
	sess, err := getAwsSession()
	if err != nil{
		fmt.Println("getAwsSession error:", err)
		return
	}

	svc := kinesis.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	

	input:= kinesis.PutRecordInput{
		Data : []byte(fmt.Sprintf("hello, message: %v", time.Now())),
		PartitionKey: aws.String("0"),
		StreamName: aws.String("test-stream"),
	}


	o, err := svc.PutRecord(&input)
	if err != nil{
		fmt.Println("PutRecord error", err)
		return
	}

	fmt.Printf("PutRecord output: %#v\n", o)
}


// GOOS=linux go build