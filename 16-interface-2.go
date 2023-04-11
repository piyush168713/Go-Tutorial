package main
import ("fmt")

func funcExample(interf interface{}) {
	val, ok1 := interf.(float64)
	fmt.Println(val, ok1)
	}
	// Defining the main method
	func main() {
		var inter interface {
	} = 98.09
	funcExample(inter)
	var intrf interface {
	} = "TesteInterface"
	funcExample(intrf)
}