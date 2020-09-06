package main

import "fmt"

func main() {
	//Initialize simple Array
	var a [5]int
	fmt.Println("emp:", a)

	//Set a value in an index of array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	//array declaration
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		c := [5]int{1, 2, 3, 4, 5}
		d := [5]int{6, 7, 8, 9, 10}
		var e [5]int
		i := 0
		for range e {
			e[i] = c[i] + d[i]
			fmt.Println(e[i])
			i++
		}
	*/
}
