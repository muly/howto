package main

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kaperys/awslocal"
)

func TestExampleLocalstack2(t *testing.T) {
	session, err := session.NewSession(aws.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	cfg := aws.NewConfig()


	// Wrap the service cfg for use with LocalStack
	awslocal.Wrap(cfg, awslocal.ServiceS3)

	svc := s3.New(session, cfg)
	log.Println(svc.ListBuckets(&s3.ListBucketsInput{}))

	// firehoseClinet = firehose.New(session, cfg)

	// myTest(t)


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

func myTest2(t *testing.T) {
	err := Handler(context.TODO(), events.KinesisEvent{})
	if err != nil {
		t.Errorf("Handler error,: %v", err)
	}




}


// 
// An error occurred (UnrecognizedClientException) when calling the ListDeliveryStreams operation: The security token included in the request is invalid.



// An error occurred (UnrecognizedClientException) when calling the ListDeliveryStreams operation: The security token included in the request is invalid.