package main

import "fmt"

func main() {
	// Memasukkan nilai N langsung (contoh N = 10)
	N := 10

	// Menghitung kode hari berikutnya (Kamis = 4)
	hariBerikutnya := (4 + N) % 7

	// Jika sisa bagi 0, maka itu adalah hari Minggu (kode 7)
	if hariBerikutnya == 0 {
		hariBerikutnya = 7
	}

	// Menampilkan hasil
	fmt.Println(hariBerikutnya)
}