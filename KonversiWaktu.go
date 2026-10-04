package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukkan jumlah detik: ")
	fmt.Scan(&N)

	jam := N / 3600
	menit := (N % 3600) / 60
	detik := N % 60

	fmt.Printf("%d %d %d\n", jam, menit, detik)
}