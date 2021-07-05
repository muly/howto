// demonstrate a simple lambda function
// inspired from https://www.youtube.com/watch?v=x_yCX4kSchY&t=331s

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rakyll/statik/fs"

	_ "github.com/muly/howto/golang/cloud/aws/lambda/static-hosting-lambda/statik" // statik files
)

// type

// type Resp struct {
// 	Message string `json:"message"`
// 	Ok      bool   `json:"ok"`
// }

//Note: if the lambda is triggered via API Gateway then the response needs to be in below format, otherwise, the request would fail eventhough the Lambda actually succeeded.
// ref https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/

func Handler(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
	filename := strings.TrimPrefix(req.Path, "/static-file-server")
	if filename == "" {
		filename = "/index.html"
	}
	log.Printf("request: %+v", req)
	log.Printf("serving file: %s", filename)
	data, err := getFile(filename)
	if err != nil {
		log.Printf("getFile() error: %v", err)
		return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusInternalServerError,
			Body:            err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		Headers:         map[string]string{"Content-Type": "text/html"},
		StatusCode:      http.StatusOK,
		Body:            string(data),
	}, nil
}

func getFile(filename string) ([]byte, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	f, err := statikFS.Open(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func main() {
	lambda.Start(Handler)
	// filename:= "" //  "/swagger.json"//""//"/index.html"
	// r, err := Handler(events.APIGatewayProxyRequest{Path: filename})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(r)
}

//// generate statik file
// statik -src=/Users/muly/go/src/github.com/muly/howto/golang/cloud/aws/lambda/static-hosting-lambda/swagerui/dist

// run go code
// go run main.go

/* //// DEPLOYMENT STEPS ////
0. build for linux and zip the executable
	- GOOS=linux go build -o lambda main.go
1. login to aws
2. from IAM roles page: https://console.aws.amazon.com/iam/home?#/roles
	- "create role" button
	- "AWS Service" as type
	- for "Choose the service that will use this role" choose "Lambda". then click Next.
	- on permissions page, make sure to add "CloudWatchLogsReadOnlyAccess" permission
	- on Review page, name the role, and then create the role.
	- Note: this role will be used by the below steps while creating the lamda function
3. go to lamda functions page: https://console.aws.amazon.com/lambda/home
	- click on "Create function" button
	- Choose "Author from scratch" template
	- fill the fields.
		-- choose Golang runtime
		-- for Role, choose the existing role, created as part of the previous steps
	- click "Create Function" button. you have technically created the lamda function, you now need to configure it)
4. configure the lamda function
	- under "Function code" section:
		-- choose the "Upload a .zip file" type
		-- choose Go runtime
		-- for "Handler" field, type the name of the executable in the zip file uploaded
		-- using "Upload" button, upload the binary zip file generated in previous steps.
	- using "Designer" section,
		-- add "API Gateway" trigger,
		-- configure it:
			--- choose "Create a new API"
			--- choose the desired security. for testing you can choose "Open" security
			--- click "Add"
	- "save" changes
5. Test
	- find the endpoint/api trigger url
	- using your favorite rest client, hit the api endpoint
	- make sure to pass the json of id, value (see go code)
		{
		"id": 1,
		"value": "test1"
		}
	- TODO: tested from asw sucessfully, but need to test gtom rest client
*/

/* //// POSSIBLE ERRORS ////
here are some possible errors, that i have seen

// when uploading the binary in aws console, if the name mentioned in the "Handler" is not same as that of binary built (in the zip file being uploaded)
{
  "errorMessage": "fork/exec /var/task/lambda: no such file or directory",
  "errorType": "PathError"
}
// same error as above is possible when the uploaded zip file has a folder within
// when not build for linux. if NOT done like this GOOS=linux go build
{
       "errorMessage":"fork/exec /var/task/lambda: exec format error",
       "errorType":"PathError"
}
// when used by the API-gateway to trigger, and the response is not in the required format. ref https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/
	Execution failed due to configuration error: Malformed Lambda proxy response

*/

/*
references:
[] https://www.youtube.com/watch?v=x_yCX4kSchY&t=331s
[] https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/
*/
