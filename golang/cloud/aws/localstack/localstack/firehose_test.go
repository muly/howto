package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	dockertest "github.com/ory/dockertest"
)

const (
	bucketName = "test-bucket-for-firehose"
	roleName   = "test-role"
)

func TestMain(m *testing.M) {
	//
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("localstack/localstack", "latest", []string{"SERVICES=s3,iam,firehose"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	endpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("4566/tcp"))
	log.Println("endpoint: ", endpoint)

	time.Sleep(30 * time.Second) // or poll for the status

	// aws session
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "empty", "test-token"),
		DisableSSL:  aws.Bool(true),
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
	})
	if err != nil {
		log.Fatalf("session.NewSession() error: %v", err)
	}

	// setup s3 bucket
	s3Client := s3.New(sess, &aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true), // Required: otherwise running into an error
	})

	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket:           aws.String(bucketName),
		GrantFullControl: aws.String("write"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(region),
		},
	})
	if err != nil {
		log.Fatalf("Could not CreateBucket: %s", err)
	}

	// setup iam role
	iamClient := iam.New(sess)

	roleResponse, err := iamClient.CreateRole(&iam.CreateRoleInput{
		RoleName:                 aws.String(roleName),
		AssumeRolePolicyDocument: aws.String("AmazonKinesisFirehoseFullAccess"),
	})
	if err != nil {
		log.Fatalf("Could not CreateBucket: %s", err)
	}



	// setup  firehose
	firehoseClinet = firehose.New(sess)

	_, err = firehoseClinet.CreateDeliveryStream(&firehose.CreateDeliveryStreamInput{
		DeliveryStreamName: aws.String(firehoseName),
		DeliveryStreamType: aws.String("DirectPut"),
		ExtendedS3DestinationConfiguration: &firehose.ExtendedS3DestinationConfiguration{
			BucketARN: aws.String(fmt.Sprintf("arn:aws:s3:%s::%s", region, bucketName)),
			RoleARN:   roleResponse.Role.Arn,
			DataFormatConversionConfiguration: &firehose.DataFormatConversionConfiguration{
				Enabled: aws.Bool(true),
				InputFormatConfiguration: &firehose.InputFormatConfiguration{
					Deserializer: &firehose.Deserializer{
						OpenXJsonSerDe: &firehose.OpenXJsonSerDe{
							// TODO
						},
					},
				},
				OutputFormatConfiguration: &firehose.OutputFormatConfiguration{
					Serializer: &firehose.Serializer{
						ParquetSerDe: &firehose.ParquetSerDe{
							// TODO 
						},
					},
				} ,
				SchemaConfiguration: &firehose.SchemaConfiguration{

				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Could not CreateDeliveryStream: %s", err)
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
		t.Errorf("Handler error: %v", err)
	}

}

/*
currently getting the below error:
	Internal Server Error
	The server encountered an internal error and was unable to complete your request. Either the server is overloaded or there is an error in the  application
*/

/*
related aws commands
	awslocal s3api list-buckets
	awslocal s3api list-objects --bucket test-bucket-for-firehose
	awslocal firehose list-delivery-streams
*/