// demonstrate a simple lambda function
// inspired from https://www.youtube.com/watch?v=x_yCX4kSchY&t=331s

package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Req struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// type Resp struct {
// 	Message string `json:"message"`
// 	Ok      bool   `json:"ok"`
// }

//Note: if the lambda is triggered via API Gateway then the response needs to be in below format, otherwise, the request would fail eventhough the Lambda actually succeeded.
// ref https://aws.amazon.com/premiumsupport/knowledge-center/malformed-502-api-gateway/
type GateWayResp struct {
	IsBase64Encoded bool `json:"isBase64Encoded"`
	StatusCode      int  `json:"statusCode"`
	//Headers `json:"headers"`
	Body string `json:"body"`
}

func Handler(req Req) (resp GateWayResp, err error) {
	return GateWayResp{
		IsBase64Encoded: false,
		StatusCode:      http.StatusOK,
		Body:            fmt.Sprintf("processing the request ID %v", req.ID),
	}, nil
	//return Resp{
	//Message: fmt.Sprintf("processing the request ID %v", req.ID),
	//Ok:      true,
	//}, nil
}

func main() {
	lambda.Start(Handler)
	// r, err:= Handler(Req{ID:1, Value:"hello"})
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(r)

}

/* some possible errors, that i have seen

// when uploading the binary in aws console, if the name mentioned in the "Handler" is not same as that of binary built (in the zip file being uploaded)
{
  "errorMessage": "fork/exec /var/task/lambda: no such file or directory",
  "errorType": "PathError"
}
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
