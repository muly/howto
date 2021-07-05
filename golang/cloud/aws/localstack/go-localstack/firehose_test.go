package main

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	// "github.com/aws/aws-sdk-go/service/s3"
	"github.com/elgohr/go-localstack"
)

func TestExampleLocalstack(t *testing.T) {
	l, err := localstack.NewInstance()
	if err != nil {
		log.Fatalf("Could not connect to Docker %v", err)
	}
	if err := l.Start(); err != nil {
		log.Fatalf("Could not start localstack %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "empty", "test-token"),
		DisableSSL:  aws.Bool(true),
		Region:      aws.String(endpoints.UsWest1RegionID),
		Endpoint:    aws.String(l.Endpoint(localstack.Firehose)),
	})
	if err != nil{
		t.Errorf("session.NewSession() error: %v", err)
		return
	}

	log.Print(l.Endpoint(localstack.Firehose))



	firehoseClinet = firehose.New(sess)

	myTest(t)


	// s3Client := s3.New(session.New(), &aws.Config{
	// 	Credentials: credentials.NewStaticCredentials("not", "empty", "test-token"),
	// 	DisableSSL:  aws.Bool(true),
	// 	Region:      aws.String(endpoints.UsWest1RegionID),
	// 	Endpoint:    aws.String(l.Endpoint(localstack.S3)),
	// })
	
	// input := &s3.GetObjectInput{
	// 	Bucket: aws.String("crudapi"),
	// 	Key:    aws.String(path),
	// }
	// s3Client.GetObject(input)

	// if err := l.Stop(); err != nil {
	// 	log.Fatalf("Could not stop localstack %v", err)
	// }
}

func myTest(t *testing.T) {
	err := Handler(context.TODO(), events.KinesisEvent{})
	if err != nil {
		t.Errorf("Handler error,: %v", err)
	}




}


// 
// An error occurred (UnrecognizedClientException) when calling the ListDeliveryStreams operation: The security token included in the request is invalid.



// An error occurred (UnrecognizedClientException) when calling the ListDeliveryStreams operation: The security token included in the request is invalid.