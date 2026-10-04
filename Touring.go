package main

import "fmt"

func main() {
	// [DEKLARASI VARIABEL]
	var kecepatan, menit int // MASUKKAN TIPE DATA
	var jam, total_jarak int // MASUKKAN TIPE DATA

	// [INPUT]
fmt.Print("Masukkan kecepatan: ") // Tambahkan baris ini
fmt.Scan(&kecepatan)

	// [PROSES]
	total_jarak = 100 + 60 + 170   // MASUKKAN TOTAL JARAK TEMPUH (100+60+170)
	menit = (total_jarak * 60) / kecepatan // BERISI TOTAL JARAK
	jam = menit / 60               // MASUKKAN RUMUS MENCARI JAM
	menit = menit % 60             // MASUKKAN SISA MENIT

	// [OUTPUT]
	fmt.Println(jam, menit) // MASUKKAN VARIABEL YANG TELAH DIPROSES
}