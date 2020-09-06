package main

import "fmt"

func main() {

	//initialize empty slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//getter and setter
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//length builtin function (just like arrays)
	fmt.Println("len:", len(s))

	//append builtin function
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//copy the slice-to-slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	//exclude element-5
	l = s[:5]
	fmt.Println("sl2:", l)

	//slices up from element-2
	l = s[2:]
	fmt.Println("sl3:", l)

	//single line declaration
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//multidimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		f := []int{1, 2, 3, 4, 5}
		g := []int{6, 7, 8, 9, 10}
		var e []int
		i := 0
		for range f {
			e = append(e, f[i]+g[i])
			fmt.Println(e[i])
			i++
		}
	*/

}
