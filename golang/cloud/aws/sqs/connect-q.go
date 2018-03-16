// example to demonstrate how to connect to an aws sqs by the queue name


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
	region := os.Getenv("AWSREGION") 
	queueName := os.Getenv("") 
 //TODO the rest 
}
