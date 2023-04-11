package main
import "fmt"

func main(){
	arr1 := [...]int{1,2,3,4}
	var arr2 = [5]int{10,20,30,40,50}
	var cars = [4]string{"GTR", "Supra", "Demon", "Challenger"}
  	fmt.Println(cars)
  	fmt.Println(cars[0])
	arr1[2] = 20
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr2[1])

	var arr3 = [5]int{1:20,2:30}
	fmt.Println(arr3)
	fmt.Println("Length of arr2:",len(arr3))
	fmt.Println("Length of cars:",len(cars))
	fmt.Println("Capacity of cars:",cap(cars))
}