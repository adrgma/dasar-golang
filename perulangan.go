package main

import (
	"fmt"
)

func main(){
	// perulangan for
	for a:=0; a<5; a++{
		fmt.Println("Angka a", a)
	}

	//bentuk lain for
	fmt.Printf("\n\nbentuk lain for\n\n")
	var b = 0
	for b<5 {
		fmt.Println("Angka b", b)
		b++
	}

	//for tanpa argumen
	fmt.Printf("\n\nfor tanpa argumen\n\n")
	var c = 0
	for {
		fmt.Println("Angka", c)
		c++
		if c == 5 {
			break
		}
	}

	//Perulangan + Label
	fmt.Printf("\n\nfor dengan label\n\n")
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}