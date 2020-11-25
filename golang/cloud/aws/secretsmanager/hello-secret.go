// TODO: check for syntax errors

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var access = os.Getenv("AWS_ACCESS_KEY_ID")
var secret = os.Getenv("AWS_SECRET_ACCESS_KEY")
var region = os.Getenv("REGION")

func getAwsSession() (*session.Session, error) {

	token := ""
	creds := credentials.NewStaticCredentials(access, secret, token)

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess, err := session.NewSession(cfg)
	if err != nil {
		fmt.Println("session failed:", err)
		return nil, err
	}
	return sess, nil
}

func getSecret(sess *session.Session, secretName string, region string) (string, string, error) {
	// Create a Secrets Manager client
	svc := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", "", err
	}
	if result.SecretString == nil {
		return "", "", fmt.Errorf("secret string empty")
	}

	payload := struct {
		User string `json:"user,omitempty"`
		Pwd  string `json:"pwd,omitempty"`
	}{}
	if err := json.Unmarshal([]byte(*result.SecretString), &payload); err != nil {
		return "", "", err
	}
	return payload.User, payload.Pwd, nil
}

func main() {
	sess, _ := getAwsSession()

	h, b, _ := getSecret(sess, "my-secrets", region)
	fmt.Println(h, b)
}