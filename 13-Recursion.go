package main
import ("fmt")

func testCount(a int) int {
	if(a == 11){
		return 0
	}
	fmt.Println(a)
	return testCount(a+1)
}

func factorial(x int) int {
	if(x == 1){
		return 1
	}
	return (x * factorial(x-1))
}

func main(){
	// testCount(1);
	fmt.Println("Factorial is",factorial(3));
}