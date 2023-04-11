package main
import ("fmt")

func main(){
	/*
	var a = map[string]string{"Car": "Ford", "Animal": "Cow", "Exam": "Test"}
	b := map[string]int{"Car": 5, "Phone": 7, "Copy": 10}

	fmt.Println("a",a)
	fmt.Println("b",b)


	// Creating Maps Using Using make()Function:
	var c = make(map[string]string)
	d := make(map[string]int)

	c["A"] = "B"
	c["Apple"] = "Fruit"
	c["Animal"] = "Dog"

	d["Copy"] = 5
	d["Phone"] = 2
	d["Laptop"] = 4

	fmt.Println("c",c)
	fmt.Println("d",d)
	fmt.Println(c["Animal"])


	// Note: The make()function is the right way to create an empty map.
	var e = make(map[string]string)
	var f map[string]int

	fmt.Println(e == nil)
	fmt.Println(f == nil)
	*/

	/*
	var arr1 = [5]int{1,2,3,5,5}
	var mp = make(map[int]int)
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		mp[arr1[i]]++;
	}

	for i := 0; i < len(mp); i++ {
		fmt.Print(i)
		fmt.Println()
		fmt.Print(mp[i])
	}
	*/

	/*
	// Checking specific element in map
	var a = map[string]int{"Car": 5, "Phone": 7, "Copy": 10}

	val1, ok1 := a["Car"]
	val2, ok2 := a["Phone"]
	_, ok3 := a["Copy"]
	val4, ok4 := a["Pens"]

	fmt.Println(val1,ok1)
	fmt.Println(val2,ok2)
	fmt.Println(ok3)
	fmt.Println(val4,ok4)
	*/

	/*
	// one map reference to the other
	var a = map[string]int{"Car": 5, "Phone": 7, "Copy": 10}
	b := a

	a["Copy"] = 12
	fmt.Println(a)
	fmt.Println(b)
	*/

	
	// Iterating over maps
	// var a = map[string]int{"Car": 5, "Phone": 7, "Copy": 10}
	// for key, val := range a {
	// 	fmt.Println(key, val)
	// }
	

	
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	
	var b = make([]string, 5)
	b = append(b, "one", "two", "three", "four")

	for key, val := range a {
		fmt.Println(key, val)
	}

	for _, element := range b {
		fmt.Println(element, a[element])
	}




	



}