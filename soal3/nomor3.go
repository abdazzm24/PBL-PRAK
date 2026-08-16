package main

import "fmt"

// Function swap: Menukar nilai dua integer melalui pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Function updateSlice: Menambahkan item baru ke slice melalui pointer
func updateSlice(s *[]string, newItem string) {
	*s = append(*s, newItem)
}

// Demonstrasi Pass by Value (menerima salinan data)
func ubahNilai(val int) {
	val = 999
}

// Demonstrasi Pass by Pointer (menerima alamat memori)
func ubahNilaiPointer(val *int) {
	*val = 999
}

func main() {
	// 1. Pengujian fungsi swap
	x, y := 10, 20
	fmt.Printf("Sebelum swap : x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Setelah swap : x = %d, y = %d\n\n", x, y)

	// 2. Pengujian fungsi updateSlice
	daftarBuah := []string{"Apel", "Pisang"}
	fmt.Println("Slice sebelum update:", daftarBuah)
	updateSlice(&daftarBuah, "Mangga")
	fmt.Println("Slice setelah update:", daftarBuah)

	// 3. Perbandingan Pass by Value vs Pass by Pointer
	angka := 50
	fmt.Println("\n--- Perbandingan Pass by Value vs Pass by Pointer ---")
	fmt.Printf("Nilai awal variabel 'angka': %d\n", angka)

	ubahNilai(angka)
	fmt.Printf("Setelah ubahNilai (Pass by Value)   : %d\n", angka)

	ubahNilaiPointer(&angka)
	fmt.Printf("Setelah ubahNilaiPointer (Pointer) : %d\n", angka)
}