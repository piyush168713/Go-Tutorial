/*
Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

Using the []datatype{values} format
Create a slice from an array
Using the make() function

*/


package main
import "fmt"

func main(){
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	fmt.Println(myslice2[3])


	/*
	create a slice by slicing an array:

	Syntax
	var myarray = [length]datatype{values} // An array
	myslice := myarray[start:end] // A slice made from the array
	*/
	var arr1 = [6]int{10, 11, 12, 13, 14,15}
	mySlice := arr1[2:4]

	fmt.Printf("myslice = %v\n", mySlice)
    fmt.Printf("length = %d\n", len(mySlice))
    fmt.Printf("capacity = %d\n", cap(mySlice))


	/*
	The make() function can also be used to create a slice.

	Syntax
	slice_name := make([]type, length, capacity)
	if capacity not defined, it will be equal to its length
	*/

	myslice := make([]int, 5, 10)
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	
	myslice1 := make([]int, 5)
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

}