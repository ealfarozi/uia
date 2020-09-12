package main

import "fmt"

//Struct
type Vertex struct {
	Lat, Long float64
}

//declaration
var v map[string]Vertex

func main() {

	//Maps declaration (empty map)
	m := make(map[string]int)

	//Maps assignment -> set key-value
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//Get a value for a key with name[key] - an example of init and declaration
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	//deleting a key
	delete(m, "k2")
	fmt.Println("map:", m)

	//Get a value for a key with name[key]
	//an example of init and declaration including handling the error
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//Maps declaration and initialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//initialization
	v = make(map[string]Vertex)

	//manual assignment based on struct
	v["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(v["Bell Labs"])

	/*
		x := make(map[string]int)
		x["key1"] = 2
		x["key2"] = 3
		x["key3"] = 4
		x["key4"] = 5
		x["key5"] = 6

		fmt.Println("before deletion: ", x)

		delete(x, "key3")

		fmt.Println("after deletion: ", x)
	*/
}
