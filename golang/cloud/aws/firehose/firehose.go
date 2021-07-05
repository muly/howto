package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/go-kit/kit/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

var firehoseClinet *firehose.Firehose

const region = "us-east-1"

func Init() {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	
	firehoseClinet = firehose.New(session.New(), cfg) //TODO: session.New() is deprecated
}


func Handler(ctx context.Context, data events.KinesisEvent ) error {
	var logger log.Logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	lctx, ok := lambdacontext.FromContext(ctx)
	if ok{
		logger = log.With(logger, "id", lctx.AwsRequestID)
	}

	fmt.Printf("count: %v\n", len(data.Records))
	logger.Log("count", len(data.Records))


	ProcessRecord(ctx, firehoseClinet)

	fmt.Print("Handler Done")
	
	return nil
}


func ProcessRecord(ctx context.Context, firehoseClinet firehoseiface.FirehoseAPI) error{
	input := firehose.PutRecordInput{
		DeliveryStreamName: aws.String("test-hose1"),
		Record: &firehose.Record{
			Data : []byte(fmt.Sprintf("hello firehose: %v",time.Now() )),
		},
	}

	_, err:= firehoseClinet.PutRecordWithContext(ctx,&input)
	if err != nil{
		fmt.Println("firehoseClinet.PutRecordWithContext error", err)
		return err
	}
	return nil
}

func main() {
	Init()

	// lambda.Start(Handler)

	err := Handler(context.TODO(), events.KinesisEvent{})
	if err != nil{
		fmt.Printf("Handler error,: %v", err)
	}

}