package main

import (
	"fmt"
)

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	// Membaca input n dan k
	fmt.Scan(&n, &k)

	// Memanggil prosedur isiArray untuk mengisi data
	isiArray(n)

	// Memanggil fungsi posisi untuk mencari posisi k
	hasil := posisi(n, k)

	// Menampilkan hasil pencarian
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	// Mengisi array data dengan n elemen yang diterima dari input
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	// Implementasi pencarian biner (binary search)
	kiri, kanan := 0, n-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah] == k {
			return tengah // Jika k ditemukan, kembalikan indeksnya
		} else if data[tengah] < k {
			kiri = tengah + 1 // Cari di sebelah kanan
		} else {
			kanan = tengah - 1 // Cari di sebelah kiri
		}
	}
	return -1 // Jika k tidak ditemukan, kembalikan -1
}
