// demonstrate triggering lambda when a file in S3 bucket is added/updated and
// read the file and make a copy into a different bucket/folder

package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// for different event type see README in https://github.com/aws/aws-lambda-go/tree/master/events
func s3Handler(ctx context.Context, s3Event events.S3Event) {
	sess := session.Must(session.NewSession())

	// get file changes summary
	output := []string{}
	for i := range s3Event.Records {
		output = append(output, fmt.Sprintf("key:%s\tevent-name:%s\tevent-source:%s\tevent-time:%s\t",
			s3Event.Records[i].S3.Object.Key,
			s3Event.Records[i].EventName,
			s3Event.Records[i].EventSource,
			s3Event.Records[i].EventTime))

		// read file
		buff := &aws.WriteAtBuffer{}
		downloader := s3manager.NewDownloader(sess)
		downloader.Download(buff, &s3.GetObjectInput{
			Bucket: aws.String(s3Event.Records[i].S3.Bucket.Name),
			Key:    aws.String(s3Event.Records[i].S3.Object.Key),
		})

		// save file to backup folder
		uploader := s3manager.NewUploader(sess)
		uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("test-muly-123"),
			Key:    aws.String(fmt.Sprintf("backup/%s", s3Event.Records[i].S3.Object.Key)),
			Body:   bytes.NewReader(buff.Bytes()),
		})
	}

	// save the file changes summary to s3
	uploader := s3manager.NewUploader(sess)
	u := uuid.New()
	key := fmt.Sprintf("responses/%s", u.String())
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("test-muly-123"),
		Key:    aws.String(key),
		Body:   bytes.NewReader([]byte(fmt.Sprintf("%v", output))),
	})
	if err != nil {
		log.Fatalf("error saving file to s3: %v", err)
	}
	log.Printf("saved the report to: %s", key)
}

func main() {
	lambda.Start(s3Handler)
}

/* ///// DEPLOYMENT STEPS ////
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o lamda

zip the binary

create lambda function

upload the zip as lambda source code

set lamda (binary name, used with the go build) as Handler in the Runtime settings.

lambda -> Roles: make sure the corresponding role has write permissions on the s3

create a bucket

add s3 as input/trigger. while creating ...
	- bi=ucket: use the above bucket while creating the trigger.
	- event type: All objects create events
	- Prefix: add prefix to indicate the folder structure
	- Suffix: add suffix to indicate the file extension (*.jpg)


to test:
	add a file/replace a file in s3, and check the cloudwatch logs to verify if the correct file info is logged


*/
