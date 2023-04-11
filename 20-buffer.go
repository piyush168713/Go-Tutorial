package main
import (
	"fmt"
	"bytes"
)

func main(){
	var strBuffer1 bytes.Buffer
	strBuffer1.WriteString("Piyush")
	strBuffer1.WriteString("Kumar")
	fmt.Println("The name is: ", strBuffer1.String());

}