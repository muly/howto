package main

import (
	"github.com/go-openapi/loads"

	"github.com/muly/howto/golang/code-gen/swagger/hello-swagger/api/handlers"
	"github.com/muly/howto/golang/code-gen/swagger/hello-swagger/gen/restapi"
	"github.com/muly/howto/golang/code-gen/swagger/hello-swagger/gen/restapi/operations"
)

func main() {

	// 1: specs
	specs, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	// 2: initialize the api server
	api := operations.NewHelloAPI(specs)
	// TODO: Add logger, middleware etc

	// 3. add handlers
	api.GetStudentByIDHandler = handlers.NewGetStudentByID()
	// api.GetStudentsHandler = // TODO: add it after implementing the this handler in ./api/handlers/get_students.go

	// 4. add handler and middleware layers
	handler := api.Serve(nil)
	// add middleware layers to handler

	// 5. add server, its config
	server := restapi.NewServer(api)
	server.Port = 8080
	server.SetHandler(handler)
	defer server.Shutdown()

	// 6. start the server
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
