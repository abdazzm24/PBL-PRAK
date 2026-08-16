package main

import "fmt"

// Struct Student
type Student struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Grade    float64 `json:"grade"`
	IsActive bool    `json:"is_active"`
}

// Value Receiver: Hanya membaca data tanpa mengubah atribut struct
func (s Student) GetInfo() string {
	return fmt.Sprintf("ID: %d | Nama: %s | IPK: %.2f | Status Aktif: %t", s.ID, s.Name, s.Grade, s.IsActive)
}

// Pointer Receiver: Memperbarui atribut Grade pada struct asli
func (s *Student) UpdateGrade(grade float64) {
	s.Grade = grade
}

// Pointer Receiver: Mengubah atribut IsActive menjadi true
func (s *Student) Activate() {
	s.IsActive = true
}

// Pointer Receiver: Mengubah atribut IsActive menjadi false
func (s *Student) Deactivate() {
	s.IsActive = false
}

func main() {
	// Inisialisasi struct
	mhs := Student{
		ID:       101,
		Name:     "Abdullah Azzam",
		Grade:    3.20,
		IsActive: false,
	}

	fmt.Println("--- Kondisi Awal ---")
	fmt.Println(mhs.GetInfo())

	fmt.Println("\n--- Mengaktifkan Status & Memperbarui Nilai ---")
	mhs.Activate()
	mhs.UpdateGrade(3.85)
	fmt.Println(mhs.GetInfo())

	fmt.Println("\n--- Non-aktifkan Status ---")
	mhs.Deactivate()
	fmt.Println(mhs.GetInfo())
}