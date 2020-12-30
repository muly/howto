
using consumer group
- to process the messages
- marking the processed messages
- restarting processing the messages by changing the group name


very good kafka consumer group ref: https://ednsquare.com/story/how-to-create-consumer-group-consumer-in-golang-with-kafka------JrFWqc

notes: as demonstrated in the code in the above blogpost, the message is received in the `ConsumeClaim()` method, and any custom code to work on the received message should be written there

