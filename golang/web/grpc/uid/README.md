UID Generator:
    GRPC 64 bit unique id generator of the below pattern. saves the generated uids to configured mongodb table with the generated uid into the _id column, which is the key column.
        1. The 20 most significant bits are all 0
        2. The next 32 most significant bits are the Unix timestamp of when the id was generated
        3. The least significant 12 bits increase monotonically within any given second, but reset to 0’s with each new Unix timestamp



-- steps: 
    Use the below steps to start the UID generator GRPC service and then run the client code to test it
    1) start mongodb container
        docker-compose up
    2) build docker image for uid server
        docker build -t uidimage .
    3) run the docker image for uid server 
        docker run -it -p 5050:5050 --network uid_mongonetwork  uidimage
    4) run the client code
        go run client/client.go 




-- generate proto go code: if any changes are made to the proto, then runthe below to regenerate the go code 
    protoc --go_out=plugins=grpc:.  proto/*.proto



-- configuration
    below are the configuration values defined as constants in code:
    1) mongoHostUrl: the url of the mongodb to connect to
    2) mongoDb: the database name to save the generated unique ids
	3) mongoTable: the collection name to save the generated unique ids
    4) wait: wait time before checking if the new second has started. used in the polling code that waits for next seconf if the max IDs are generated for that given second.
        Note: 4095 is the max num corresponding to the 12 bits allocated for the sequence number for uid to be used within the second.
    5) qty: on the client side, this constant indicates the number of uids to be generated


external references:
    https://www.melvinvivas.com/my-first-go-microservice/
    


roadmap:
    1) need to handle TODOs recorded in the code
    2) larger number of IDs requested will result in longer execution time: this is because we are using the unix timestamp in the id of when the id is generated. also we only have 12 bits to maintain the counter of IDs generated within a second. that limits the number of uid generated to 4096 per minute, and so larger request means longer execution time. need a better approach
    
