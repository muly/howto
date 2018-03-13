// demonstrate a simple lambda function
// inspired from https://www.youtube.com/watch?v=x_yCX4kSchY&t=331s

package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Req struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Resp struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(req Req) (resp Resp, err error) {
	return Resp{
		Message: fmt.Sprintf("processing the request ID %v", req.ID),
		Ok:      true,
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
