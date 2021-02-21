package main

import (
	"fmt"
	"math"
)

//Vertex Struct
type Vertex struct {
	X, Y float64
}

//Abs is the function to get an absolute value from Vertex
func (v Vertex) Abs() float64 {
	fmt.Println("Abs:", v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Scale = Vertex modifiers - 1 - value receivers
//(v Vertex) = inheritance
//(f float64) = input parameter
//after (f float64) could be output parameter/response, please define the type or the struct's name
func (v Vertex) Scale(f float64) {
	fmt.Println("Scale with value")
	fmt.Println("before:", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("after:", v)
}

//Scale2 = Vertex modifiers - 2 - pointer receivers
func (v *Vertex) Scale2(f float64) {
	fmt.Println("Scale with pointer")
	fmt.Println("before:", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("after:", v)
}

//Scale3 = Vertex modifiers - 3 - value receivers
func (v Vertex) Scale3(f float64) {
	fmt.Println("Scale with value receiver")
	fmt.Println("before:", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("after:", v)
}

//Main Program
func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	v.Scale2(10)
	fmt.Println(v.Abs())
	v.Scale3(10)
	fmt.Println(v.Abs())
}
