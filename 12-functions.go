package main

import ("fmt")

func myMsg(){
	fmt.Println("Function created")
}

func namess(fname string){
	fmt.Println("hello", fname)
}

func familyName(fname string, age int){
	fmt.Println("Hello", fname, "Your age is", age)
}

func myFunc(a int, b int) int{
	return a+b
}

func myFunc1(a int, b int) (result int){
	result = a+b
	return result
}

// multi return values
func myFunc2(a int, b string) (res int, txt string) {
	res = a + a
	txt = b + "World"
	return
}



func main(){
	myMsg()
	namess("Piyush")
	familyName("Piyush", 20)
	fmt.Println(myFunc(10,20))
	fmt.Println(myFunc1(10,20))
	fmt.Println(myFunc2(5, "Hello"))
}