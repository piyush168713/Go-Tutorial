package main
import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

// channels are used in multithread context

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

func main(){
	// dataChan := make(chan int)
	dataChan := make(chan int,)

	go func(){
		wg := sync.WaitGroup{}
		for i:= 0; i < 1000; i++ {
			wg.Add(1)
			go func(){
				defer wg.Done()
				result := doWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Println("n = ",n)
	}
	
}