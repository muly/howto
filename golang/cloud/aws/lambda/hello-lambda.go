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

type GateWayResp struct { //Note: if the lambda is triggered via API Gateway then the response needs to be in below format, otherwise, the request would fail eventhough the Lambda actually succeeded
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

		//Message: fmt.Sprintf("processing the request ID %v", req.ID),
		//Ok:      true,
	}, nil
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
