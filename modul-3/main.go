package main

import (
	"fmt"
	"math"
)

//Constant dengan menggunakan tipe data
const s string = "constant"

func main() {
	fmt.Println(s)

	//Constant untuk tipe data int
	const n = 500000000

	//Constant untuk tipe data int dengan eksponen
	const d = 3e20 / n
	fmt.Println(d)

	//Constant untuk tipe data int64
	fmt.Println(int64(d))

	//Constant dengan fungsi math sin
	fmt.Println(math.Sin(n))
}
