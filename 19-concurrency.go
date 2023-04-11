package main
import (
	"fmt"
	"time"
)

/*
func show (name string){
	for i:= 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name)
	}
}

func main(){
	go show("Ajay")
	show("Ranjan")
}
*/

func main(){
	fmt.Println("Hello")
	go func(){
		fmt.Println("heyyyy")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("we have done")
}