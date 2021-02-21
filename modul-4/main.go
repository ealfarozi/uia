package main

import "fmt"

func main() {
	//classic For
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//For with single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//For without condition
	for {
		fmt.Println("loop")
		break
	}

	//classic for with continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//latihan
	// var m, k = 0, 0
	// for {
	// 	m = m + k
	// 	if m > 5000000000 {
	// 		fmt.Println("loop terakhir: ", k)
	// 		fmt.Println("jumlah nilai: ", m)
	// 		break
	// 	}
	// 	k++
	// }

}
