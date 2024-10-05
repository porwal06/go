package main

import "time"

func main() {
	// Create channel for getting status for each function
	channel := make(chan bool)
	//var channel chan bool

	// Function to print text
	println("This is main function")

	// Added go keyword to run the function concurrently.
	go printSlowText("This is 1st text", channel)
	go printText("This is 2nd text", channel)
	go printSlowText("This is slow text", channel)
	go printText("This is 4th text", channel)

	// Wait for all the functions to complete
	println(<-channel) // Print channel value
	<-channel
	<-channel
	<-channel
}

func printText(text string, channel chan bool) {
	println(text)
	channel <- true
}

func printSlowText(text string, channel chan bool) {
	// Simulate a slow operation
	time.Sleep(5 * time.Second)
	println(text)
	channel <- true
}
