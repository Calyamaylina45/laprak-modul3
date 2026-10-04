package main

import "fmt"

func main() {
	// [DEKLARASI VARIABEL]
	var gajiPokok, jamLembur int
	var bonusLembur, potongan, gajiBersih float64

	// [INPUT]
	fmt.Scan(&gajiPokok, &jamLembur)

	// [PROSES]
	bonusLembur = float64(jamLembur) * 45000.0
	potongan = (0.02 + 0.035) * float64(gajiPokok) // 2% asuransi + 3.5% dana pensiun = 5.5%
	gajiBersih = float64(gajiPokok) + bonusLembur - potongan

	// [OUTPUT]
	fmt.Println(int(gajiBersih)) // Konversi ke int agar sesuai contoh keluaran bulat
}