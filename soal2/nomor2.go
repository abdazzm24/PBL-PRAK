package main

import "fmt"

func main() {
	// Lima variabel dengan tipe berbeda
	var nama string = "Azzam"
	var umur int = 20
	var ipk float64 = 3.85
	var aktif bool = true
	var mataKuliah []string = []string{
		"Pemrograman Backend Lanjut",
		"Keamanan Cyber",
		"Basis Data",
	}

	fmt.Println("=== DATA VARIABEL ===")
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("IPK:", ipk)
	fmt.Println("Aktif:", aktif)
	fmt.Println("Mata Kuliah:", mataKuliah)

	// Map untuk menyimpan data mahasiswa
	// Nama sebagai key dan nilai sebagai value
	nilaiMahasiswa := make(map[string]int)

	// Menambahkan data ke map
	nilaiMahasiswa["Azzam"] = 90
	nilaiMahasiswa["Diaul"] = 85
	nilaiMahasiswa["Rehan"] = 95

	fmt.Println("\n=== DATA MAHASISWA ===")
	fmt.Println(nilaiMahasiswa)

	// Membaca data dengan pengecekan keberadaan
	nilai, ada := nilaiMahasiswa["Raihan"]

	if ada {
		fmt.Println("Nilai Raihan:", nilai)
	} else {
		fmt.Println("Data Raihan tidak ditemukan")
	}

	// Mengecek mahasiswa yang tidak ada
	nilai, ada = nilaiMahasiswa["Andi"]

	if ada {
		fmt.Println("Nilai Andi:", nilai)
	} else {
		fmt.Println("Data Andi tidak ditemukan")
	}

	// Menghapus data dari map
	delete(nilaiMahasiswa, "Diaul")

	fmt.Println("\n=== SETELAH DIAUL DIHAPUS ===")
	fmt.Println(nilaiMahasiswa)

	// Menelusuri seluruh isi map
	fmt.Println("\n=== SELURUH DATA MAHASISWA ===")

	for nama, nilai := range nilaiMahasiswa {
		fmt.Printf("Nama: %s, Nilai: %d\n", nama, nilai)
	}
}