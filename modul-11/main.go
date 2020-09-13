package main

import (
	"fmt"
	"math"
)

//Vertex Struct
type Vertex struct {
	X, Y float64
}

//get an absolute value from Vertex
func (v Vertex) Abs() float64 {
	fmt.Println("scale in:", v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Vertex modifiers - 1 - value receivers
func (v Vertex) Scale(f float64) {
	fmt.Println("before:", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("after:", v)
}

//Vertex modifiers - 2 - pointer receivers
func (v *Vertex) Scale_2(f float64) {
	fmt.Println("before:", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("after:", v)
}

//Vertex modifiers - 3 - value receivers
func (v Vertex) Scale_3(f float64) {
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
	v.Scale_2(10)
	fmt.Println(v.Abs())
	v.Scale_3(10)
	fmt.Println(v.Abs())

}
