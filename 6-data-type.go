package main
import "fmt"


// Type     	Size	     Range
// float32   	32 bits  	-3.4e+38 to 3.4e+38.
// float64   	64 bits  	-1.7e+308 to +1.7e+308.

func main(){
	var a int = -10
	var b string = "piyush"
	var c bool = true
	var d float64 = 5.44
	var f float64 = 1.7e+308

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	var e uint = 10
	fmt.Println(e)
}


// Signed integers - can store both positive and negative values
// Unsigned integers - can only store non-negative values

