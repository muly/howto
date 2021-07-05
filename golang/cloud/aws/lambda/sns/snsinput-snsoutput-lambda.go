package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func handler(ctx context.Context, snsEvent events.SNSEvent)  events.SNSEvent{

	snsSvc := sns.New(session.New())

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		_, err:= snsSvc.Publish(&sns.PublishInput{
			
		})
		if err != nil{

		}
		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
	}
}
