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

	// Menemukan calon dengan suara terbanyak dan suara terbanyak kedua
	var max1, max2 int
	var calon1, calon2 int

	// Menyaring calon dengan suara terbanyak pertama dan kedua
	for i := 1; i <= 20; i++ {
		if suara[i] > max1 {
			max2 = max1
			calon2 = calon1
			max1 = suara[i]
			calon1 = i
		} else if suara[i] > max2 {
			max2 = suara[i]
			calon2 = i
		}
	}

	// Menampilkan hasil ketua dan wakil ketua RT
	fmt.Printf("Ketua RT: %d\n", calon1)
	fmt.Printf("Wakil ketua: %d\n", calon2)
}
