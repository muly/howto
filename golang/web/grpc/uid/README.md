

-- generate code
protoc --go_out=plugins=grpc:.  *.proto


-- run server
go run main.go uid.pb.go 

-- run client 
go run client.go uid.pb.go