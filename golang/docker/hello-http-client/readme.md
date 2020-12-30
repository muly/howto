// demonstrating how the go http client run inside a docker container could connect to a web application running on the host (as localhost)


// 1: run the server
    docker build -t sample-server sample-server/.

    docker run --rm -p 8080:8080 sample-server
    


// 2: run the client's docker container with --network option
    // 
    // docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
    // 		docker build -t hello-client client/.
    // docker run to create the container from the specified image (using tag in below example)
    // 		docker run --rm --network="host"  hello-client





troubleshooting:
ERROR 1: client.Do error: Get http://127.0.0.1:8080: dial tcp 127.0.0.1:8080: getsockopt: connection refused
    ref: https://dev.to/bufferings/access-host-from-a-docker-container-4099
    solution: Note: server has to be run using the docker container instead of runnning directly using go run like `go run sample-server/server.go`. if not the host network will not be accessible as discussed in the above link.

