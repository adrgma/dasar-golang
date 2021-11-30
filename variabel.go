package main

import "fmt"

func main(){
	//cara deklarasi variabel 1
	var namaDepan string = "Aditia Arga"
	
	//cara deklarasi variabel 2
	var namaBelakang string
	namaBelakang = "Pratama"
	fmt.Printf("Halo %s %s!\n", namaDepan, namaBelakang)

	//Cara deklarasi variabel 3
	//deklarasi berdasarkan tipe data nilai variabel
	// saat deklarasi harus menggunakan ":=" selanjutnya cukup "="
	salam := "Hai"
	fmt.Printf("%s %s %s! ini salam kedua\n", salam, namaDepan, namaBelakang)
	salam = "hello"
	fmt.Printf("%s %s %s! ini salam ketiga\n", salam, namaDepan, namaBelakang)

	//Deklarasi multi variabel
	//hal ini bisa dikombinasikan dengan cara deklarasi 1,2, dan 3
	var first, second, third int
	first, second, third = 1,2,3
	fmt.Println(first,second,third)

	//contoh deklarasi variabel multi dengan cara 3
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	//variabel sampah
	_ = "belajar"

	//Variabel pointer
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

	//konstanta
	//digunakan jika nilai variabel sudah pasti dan tidak berubah
	const pi =3.14
	fmt.Print(pi)
}