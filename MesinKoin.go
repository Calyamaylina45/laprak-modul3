package main

import "fmt"

func main() {
	// [DEKLARASI VARIABEL]
	var koin, emas, perak, tembaga, sisa int

	// [INPUT]
	fmt.Scan(&koin)

	// [PROSES]
	emas = koin / 9    // Menghitung jumlah koin emas (1 emas = 9 koin)
	sisa = koin % 9    // Menghitung sisa koin setelah dikonversi ke emas

	perak = sisa / 3   // Menghitung jumlah koin perak dari sisa koin (1 perak = 3 koin)
	tembaga = sisa % 3 // Koin yang tersisa menjadi koin tembaga

	// [OUTPUT]
	fmt.Println(emas, perak, tembaga)
}