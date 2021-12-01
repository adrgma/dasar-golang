package main

import(
	"fmt"
)

func main(){
	var names[4] string
	names[0] = "Aditia"
	names[1] = "Arga"
	names[2] = "Pratama"
	names[3] = " adalah nama saya"

	fmt.Println(names[0],names[1],names[2],names[3])

	// inisialisasi array juga bisa seperti ini
	fmt.Printf("\n\narray langsung diisi\n\n")
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//array multidimensi
	var numbersA = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	
	fmt.Println("numbers1", numbersA)
	fmt.Println("numbers2", numbers2)
}