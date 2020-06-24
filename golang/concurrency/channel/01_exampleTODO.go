// basic channel example for syntax


package main



func main(){

	// declare a channel
	var c chan int

	// declare and initialize using make
	var c = make(chan int)


	// declare and initialize using make: short hand
	c := make(chan int)

	// c <- 1 // sending is a blocking operation if no one is ready to receive from the channel: 	fatal error: all goroutines are asleep - deadlock!

	// <- c // receiving is a blocking operation if no one is ready to send into the channel: 		fatal error: all goroutines are asleep - deadlock!

}