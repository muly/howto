package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
<<<<<<< HEAD
	"github.com/aws/aws-lambda-go/lambda"
=======
	//"github.com/aws/aws-lambda-go/lambda"
>>>>>>> 576ae10... firehose aws localstack test wip
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var firehoseClinet *firehose.Firehose

<<<<<<< HEAD
const (
	region       = "us-east-1"
	firehoseName = "test-hose1"
)

=======
>>>>>>> 576ae10... firehose aws localstack test wip
func Init() {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

<<<<<<< HEAD
	if aws_access_key_id == "" || aws_secret_access_key == "" {
		log.Fatal("aws keys empty")
	}

=======
>>>>>>> 576ae10... firehose aws localstack test wip
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
<<<<<<< HEAD
		panic(err)
	}

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

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
		log.Println("firehoseClinet.PutRecordWithContext error", err)
		return err
	}

	fmt.Print("Handler Done")

=======
		fmt.Printf("bad credentials: %s", err)
	}

	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)

	
	firehoseClinet = firehose.New(session.New(), cfg) //TODO: session.New() is deprecated
}


func Handler(ctx context.Context, data events.KinesisEvent ) error {
	input := firehose.PutRecordInput{
		DeliveryStreamName: aws.String("test-hose1"),
		Record: &firehose.Record{
			Data : []byte(fmt.Sprintf("hello firehose: %v",time.Now() )),
		},
	}

	_, err:= firehoseClinet.PutRecordWithContext(ctx,&input)
	if err != nil{
		return err
	}

	log.Print("Handler Done")
	
>>>>>>> 576ae10... firehose aws localstack test wip
	return nil
}

func main() {
	Init()

<<<<<<< HEAD
	lambda.Start(Handler)

	// // to run locally
	// err := Handler(context.TODO(), events.KinesisEvent{})
	// if err != nil {
	// 	log.Printf("Handler error,: %v", err)
	// }
}
=======
	// lambda.Start(Handler)

	err := Handler(context.TODO(), events.KinesisEvent{})
	if err != nil{
		log.Printf("Handler error,: %v", err)
	}

}
>>>>>>> 576ae10... firehose aws localstack test wip
