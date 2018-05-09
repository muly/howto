package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jung-kurt/gofpdf"
)

var svc *s3.S3
var bucketName string

func Init() {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucketName = os.Getenv("aws_bucket")

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		panic(err)
	}

	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc = s3.New(session.New(), cfg) //TODO: session.New() is deprecated
}

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	var b bytes.Buffer

	err := pdf.Output(&b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := "hello.pdf"

	path := "/myfolder/" + fileName
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		Body:   bytes.NewReader(b.Bytes()),
		//ContentLength: aws.Int64(size),
		ContentType: aws.String("application/pdf"),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))

	return

}
