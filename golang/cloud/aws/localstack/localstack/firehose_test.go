package main

import (
	"context"
	"log"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	dockertest "github.com/ory/dockertest"
)

var s3Client *s3.S3
var iamClient *iam.IAM

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("localstack/localstack", "latest", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		// aws session
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials("not", "empty", "test-token"),
			DisableSSL:  aws.Bool(true),
			Region:      aws.String(endpoints.UsWest1RegionID),
			Endpoint:    aws.String("localhost:4566"),
		})
		if err != nil {
			log.Fatalf("session.NewSession() error: %v", err)
		}

		// s3 client
		s3Client = s3.New(sess)
		firehoseClinet = firehose.New(sess)
		iamClient = iam.New(sess)

		// create s3
		buckerName := "test-bucket-for-firehose"
		_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(buckerName),
		})
		if err != nil {
			// TODO
		}
		bucketARN := fmt.Sprintf("arn:aws:s3:::%s", buckerName)

		// create role
		roleOut, err := iamClient.CreateRole(&iam.CreateRoleInput{
			RoleName:                 aws.String("test-role-for-firehose"),
			AssumeRolePolicyDocument: aws.String("some-text-shoould-be-ok"),
		})
		if err != nil {
			// TODO
		}
		roleARN := *roleOut.Role.Arn

		// create firehose
		createDeliveryStreamInput := firehose.CreateDeliveryStreamInput{
			DeliveryStreamName: aws.String("test-hose1"),
			DeliveryStreamType: aws.String("DirectPut"),
			ExtendedS3DestinationConfiguration: &firehose.ExtendedS3DestinationConfiguration{
				BucketARN: aws.String(bucketARN),
				RoleARN:   aws.String(roleARN),
			},
		}
		_, err = firehoseClinet.CreateDeliveryStream(&createDeliveryStreamInput)
		if err != nil {
			// TODO
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	err := Handler(context.TODO(), events.KinesisEvent{})
	if err != nil {
		t.Errorf("Handler error,: %v", err)
	}

}



/*
currently getting the below error:
	Internal Server Error
	The server encountered an internal error and was unable to complete your request. Either the server is overloaded or there is an error in the  application
*/