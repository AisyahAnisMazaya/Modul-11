package main

import (
	"fmt"
)

func main() {
	var suara [100]int
	var suaraMasuk, suaraSah int
	var input int

	// Membaca input hingga ditemukan angka 0
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}

		// Validasi suara, hanya angka antara 1 hingga 20 yang dianggap sah
		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
		suaraMasuk++
	}

	// Menampilkan jumlah suara yang masuk dan suara sah
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	// Menampilkan hasil suara sah per calon
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
