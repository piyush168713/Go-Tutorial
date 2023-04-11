package main
import ("fmt")

type tankDimension interface{
	getArea() float64
	getVolume() float64
}

type data struct{
	r float64
	h float64
}

func (d data) getArea() float64{
	return 2*d.r*d.h + 2*3.14*d.r*d.r
}

func (d data) getVolume() float64{
	return 3.14 * d.r * d.r * d.h
}

func main(){
	var tank tankDimension
	tank = data{10,14}
	fmt.Println("The area of tank: ", tank.getArea())
	fmt.Println("The volume of tank: ", tank.getVolume())
}