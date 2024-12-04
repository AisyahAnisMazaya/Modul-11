*NOMOR 1*
Deskripsi Program
Program di atas adalah aplikasi berbasis Go yang digunakan untuk merekapitulasi hasil perolehan suara dalam sebuah proses pemungutan suara sederhana. Program ini menerima input angka yang mewakili pilihan suara, memvalidasi suara yang sah, dan kemudian menampilkan rekapitulasi hasil pemungutan suara. Berikut adalah penjelasan lebih rinci:
Fitur Utama
Input Suara:
Program membaca input angka secara berulang hingga angka 0 dimasukkan.
Angka 0 digunakan sebagai tanda akhir input.
Validasi Suara:
Hanya suara dengan angka antara 1 hingga 20 yang dianggap sah.
Suara di luar rentang tersebut tetap dihitung sebagai suara masuk, tetapi tidak dianggap sah.
Penyimpanan Data Suara:
Suara sah dihitung per calon dan disimpan dalam array suara dengan indeks 1 hingga 20.
Rekapitulasi dan Output:
Program menampilkan jumlah suara masuk dan suara sah.
Untuk setiap calon yang mendapatkan suara sah (angka 1 hingga 20), jumlah suaranya ditampilkan.
Struktur Program
Deklarasi Variabel:
suara: Array untuk menyimpan jumlah suara setiap calon (indeks 1-20).
suaraMasuk: Menyimpan total suara yang masuk.
suaraSah: Menyimpan total suara sah.
input: Variabel sementara untuk menyimpan input angka.
Proses Utama:
Loop membaca input suara hingga ditemukan angka 0.
Validasi dilakukan untuk memastikan suara hanya dihitung sebagai sah jika dalam rentang 1-20.
Hasil Akhir:
Menampilkan total suara yang masuk (suaraMasuk).
Menampilkan total suara sah (suaraSah).
Menampilkan jumlah suara yang diperoleh setiap calon dalam rentang 1-20.

*NOMOR 2*
Deskripsi Program
Program di atas adalah aplikasi berbasis Go untuk rekapitulasi hasil pemungutan suara sederhana. Program ini tidak hanya menghitung jumlah suara yang masuk dan suara sah, tetapi juga menentukan dua calon dengan suara terbanyak untuk menjadi Ketua dan Wakil Ketua RT. Berikut adalah penjelasan lebih rinci:
Fitur Utama
Input Suara:
Program membaca angka-angka sebagai input suara.
Proses input dihentikan ketika angka 0 dimasukkan.
Validasi Suara:
Hanya suara dalam rentang 1 hingga 20 yang dianggap sah.
Suara yang tidak sah tetap dihitung sebagai suara masuk, tetapi tidak diproses lebih lanjut.
Rekapitulasi Suara:
Program menyimpan jumlah suara setiap calon dalam array suara.
Menghitung jumlah total suara masuk (suaraMasuk) dan jumlah suara sah (suaraSah).
Pemilihan Ketua dan Wakil Ketua:
Program mencari dua calon dengan suara terbanyak.
Calon dengan suara terbanyak pertama ditetapkan sebagai Ketua RT.
Calon dengan suara terbanyak kedua ditetapkan sebagai Wakil Ketua RT.
Output Hasil:
Total suara masuk dan suara sah.
Calon yang menjadi Ketua dan Wakil Ketua RT.
Struktur Program
Deklarasi Variabel:
suara: Array untuk menyimpan jumlah suara setiap calon (indeks 1-20).
suaraMasuk: Total suara yang masuk.
suaraSah: Total suara sah.
max1, max2: Jumlah suara terbanyak pertama dan kedua.
calon1, calon2: Indeks calon dengan suara terbanyak pertama dan kedua.
Proses Utama:
Input suara diterima dalam loop hingga angka 0 dimasukkan.
Validasi dilakukan untuk memastikan hanya suara sah yang dihitung ke dalam array.
Penentuan Ketua dan Wakil Ketua:
Melakukan iterasi pada array untuk menemukan calon dengan suara terbanyak pertama (max1) dan kedua (max2).
Output Hasil:
Menampilkan total suara, jumlah suara sah, Ketua RT, dan Wakil Ketua RT.

*NOMOR 3*
Deskripsi Program
Program di atas adalah implementasi pencarian elemen dalam array menggunakan pencarian biner (binary search). Program ini menerima input berupa jumlah elemen array (n), elemen yang dicari (k), dan elemen-elemen array itu sendiri. Jika elemen yang dicari ditemukan, program mengembalikan posisi indeks elemen tersebut; jika tidak, program menampilkan pesan bahwa elemen tidak ada.
Fitur Utama
Input Data:
Program menerima input jumlah elemen array (n) dan elemen yang akan dicari (k).
Array diisi dengan elemen-elemen yang diinput oleh pengguna.
Proses Pencarian:
Menggunakan algoritma pencarian biner, yang efisien untuk array yang terurut.
Program mencari elemen k di dalam array dan mengembalikan posisi indeksnya jika ditemukan.
Output Hasil:
Menampilkan indeks elemen jika ditemukan.
Menampilkan pesan TIDAK ADA jika elemen tidak ditemukan.
Struktur Program
Konstanta dan Variabel Global:
NMAX: Konstanta yang menentukan batas maksimum ukuran array (1.000.000 elemen).
data: Array global untuk menyimpan elemen-elemen input.
Prosedur dan Fungsi:
isiArray(n int):
Mengisi array data dengan elemen-elemen input hingga jumlah n.
posisi(n, k int) int:
Mengimplementasikan pencarian biner untuk mencari elemen k dalam array data yang memiliki n elemen.
Proses Utama:
Program membaca input n (jumlah elemen), k (elemen yang dicari), dan elemen-elemen array.
Memanggil posisi untuk mencari elemen k.
Menampilkan hasil pencarian.
Detail Algoritma Pencarian Biner
Kondisi Awal:
Dua indeks, kiri (awal array) dan kanan (akhir array), digunakan untuk melacak batas pencarian.
Langkah-langkah:
Hitung indeks tengah (tengah = (kiri + kanan) / 2).
Bandingkan elemen di indeks tengah dengan elemen yang dicari (k):
Jika sama, kembalikan indeks tengah.
Jika elemen tengah lebih kecil dari k, cari di bagian kanan (kiri = tengah + 1).
Jika elemen tengah lebih besar dari k, cari di bagian kiri (kanan = tengah - 1).
Ulangi hingga elemen ditemukan atau kiri > kanan.
Kompleksitas:
Waktu: 
𝑂
(
log
𝑛
)
O(logn) karena array dipangkas menjadi separuh setiap iterasi.
Ruang: 
𝑂
(
1
)
O(1) karena tidak memerlukan struktur data tambahan.

