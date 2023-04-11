package main
import "fmt"

func main(){
	// x := 10
	// y := 20

	// if(x < y){
	// 	fmt.Println(y, "is greater than", x)
	// }

	/*
	time := 20

	if(time < 18){
		fmt.Println("Good Morning")
	} else if(time > 22){
		fmt.Println("Good night")
	} else{
		fmt.Println("Good evening")
	}
	*/

	num := 13
	if(num > 10){
		fmt.Println("num is greater than 10")
		if(num > 15){
			fmt.Println("Num is greater than 15")
		}
	} else {
		fmt.Println("Num is less than 10")
	}

}