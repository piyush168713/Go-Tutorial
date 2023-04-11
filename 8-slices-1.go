package main
import ("fmt")

func main() {
	/*
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 10,11)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %v\n", len(myslice1))
  fmt.Printf("capacity = %v\n", cap(myslice1))
  */

  /*
  a := []int{1,2,3}
  b := []int{4,5,6}
  c := append(a,b...)

  fmt.Printf("Slice: %v\n", c)
  fmt.Printf("Length: %v\n", len(c))
  fmt.Printf("Capacity: %v\n", cap(c))
  */

  /*
  arr1 := []int{9,10,11,12,13,14}
  myslice2 := arr1[1:5]
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))

  myslice2 = arr1[1:3]
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))

  myslice2 = append(myslice2, 20,21,22)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
  */


  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))


}