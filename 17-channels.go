package main
import ("fmt")

// channels are used in multithread context

/*
func main(){
	// dataChan := make(chan int)
	dataChan := make(chan int, 2)

	dataChan <- 789
	dataChan <- 277
	
	n := <- dataChan
	fmt.Println("n = ",n)
	
	n = <- dataChan
	fmt.Println("n = ",n)

}
*/



/*

c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.

*/

/*
func testFunc(channel chan int){
	fmt.Println("The result of sum is: ", 123 + <-channel)   // value received from channel
}

func main(){
	fmt.Println("This is the start of main method.")
	chan1 := make(chan int)

	go testFunc(chan1)     // go func used to active goroutines or pass data to one goroutine to other
	chan1 <- 23    // passing val in channel
	fmt.Println("End of main function")
}
*/

/*
func test(channel chan string) {
	for i := 0; i < 4; i++ {
	channel <- "test-channel"
	}
	close(channel)
	}
	func main() {
	channel := make(chan string)
	go test(channel)
	for {
	response, check := <-channel
	if check == false {
	fmt.Println("The channel is close ", check)
	break
	}
	fmt.Println("Here The channel is open ", response, check)
	}
}
*/

func main(){
	channel := make(chan string)

	go func(){
		channel <- "Go"
		channel <- "Java"
		channel <- "C++"
		channel <- "Shell"
		channel <- "Bash"
		close(channel)
	}()

	for response := range channel{
		fmt.Println("Channel of the language: ", response)
	}
}