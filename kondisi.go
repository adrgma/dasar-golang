package main

import (
	"fmt"
)

func main(){
	//if-else if-else
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai penuh")
	} else if point > 5 {
		fmt.Println("lulus di atas kkm")
	} else if point == 4 {
		fmt.Println("hampir lulus tapi di bawah kkm")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//variabel temporary
	var angka = 8840.0

	if percent := angka/100; percent >= 100{
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//switch-case
	var menu = 1

	switch menu {
	case 4,2:
		fmt.Println("ohayou")
	case 3:
		fmt.Println("good morning")
	default:
		{
			fmt.Println("Selamat pagi")
			fmt.Println("hai, pagi")
		}
	}

	// switch dengan gaya if=else
	point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
	//gunakan fallthorugh jika pengecekan akan tetap dilanjutkan

	//seleksinya bersarang juga bisa
	point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}