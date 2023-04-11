// Here we are using example of interface with switch cases
package main
import "fmt"
func funcExample(intr interface{}) {
// Using type switch
	switch intr.(type) {
	case string:
	fmt.Println("\n The given type is String: ", intr.(string))
	case int:
	fmt.Println("\n The given type is Integer: int, Value:", intr.(int))
	case float64:
	fmt.Println("\n The given type is Float is : ", intr.(float64))
	default:
	fmt.Println("\n The given type is not found here")
	}
}
// Defining the main method
func main() {
	funcExample("TestExample")
	funcExample(34)
	funcExample(34.9)
	funcExample(true)
	
}