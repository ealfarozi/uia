package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var cplus, python, java bool = true, false, false
	var i, j int = 1, 2
	fmt.Println(cplus, python, java, i, j)

	var f, g, h = true, false, "no"
	var k, m = 1, 2.5

	fmt.Println(f, g, h, k, m)

	n := "apple"
	fmt.Println(n)

	p, q := "orange", "juice"
	fmt.Println(p, q)

	r, s := "guava", 2
	fmt.Println(r, s)
}
