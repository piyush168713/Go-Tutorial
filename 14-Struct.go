package main
import ("fmt")

type Person struct {
	name string
	age int
	job string
	salary int
}

func printPerson(per Person){
	fmt.Println("Name:",per.name)
	fmt.Println("Age:",per.age)
	fmt.Println("Job:",per.job)
	fmt.Println("Salary:",per.salary)
}

func main(){
	var per1 Person
	var per2 Person

	per1.name = "Piyush"
	per1.age = 20
	per1.job = "SDE"
	per1.salary = 30000

	per2.name = "Ashish"
	per2.age = 21
	per2.job = "SDE"
	per2.salary = 50000

	printPerson(per1);
	printPerson(per2);
}