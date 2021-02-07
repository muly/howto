package main

import (
	"context"
	"fmt"
	"log"
	"os"
<<<<<<< HEAD
=======
	"testing"
>>>>>>> 1c1dbad... firehose aws localstack example
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

<<<<<<< HEAD
var firehoseClinet *firehose.Firehose

const (
	region       = "us-east-1"
	firehoseName = "test-hose1"
)

func Init() {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if aws_access_key_id == "" || aws_secret_access_key == "" {
		log.Fatal("aws keys empty")
	}

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
=======
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
>>>>>>> 1c1dbad... firehose aws localstack example
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
=======
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
		},
	})
	if err != nil {
		log.Fatalf("Could not CreateDeliveryStream: %s", err)
	}
>>>>>>> 1c1dbad... firehose aws localstack example

	firehoseClinet = firehose.New(session.New(), cfg) //TODO: session.New() is deprecated
}

func Handler(ctx context.Context, data events.KinesisEvent) error {
	input := firehose.PutRecordInput{
		DeliveryStreamName: aws.String(firehoseName),
		Record: &firehose.Record{
			Data: []byte(fmt.Sprintf("hello firehose: %v", time.Now())),
		},
	}

	_, err := firehoseClinet.PutRecordWithContext(ctx, &input)
	if err != nil {
<<<<<<< HEAD
		log.Println("firehoseClinet.PutRecordWithContext error", err)
		return err
=======
		t.Errorf("Handler error: %v", err)
>>>>>>> 1c1dbad... firehose aws localstack example
	}

	fmt.Print("Handler Done")

	return nil
}

<<<<<<< HEAD
func main() {
	Init()

	lambda.Start(Handler)

	// // to run locally
	// err := Handler(context.TODO(), events.KinesisEvent{})
	// if err != nil {
	// 	log.Printf("Handler error,: %v", err)
	// }
}
=======
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
>>>>>>> 1c1dbad... firehose aws localstack example
