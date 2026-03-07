package main

import (
	"errors"
	"fmt"
)

var ( //const tidak bisa panggil fungsi errors.New()
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "eko" {
		return NotFoundError
	}

	// sukses

	return nil
}

func main() {
	err := GetById("eko")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}

// ==============================
// ERROR HANDLING DI GOLANG
// ==============================

// Di Go, error adalah VALUE, bukan exception
// Interface error:
// type error interface {
//     Error() string
// }

// ==============================
// 1. PACKAGE errors
// ==============================

// errors.New()
// Digunakan untuk membuat error sederhana dengan pesan statis
// Contoh:
// return errors.New("terjadi kesalahan")

// ==============================
// 2. fmt.Errorf()
// ==============================

// Digunakan untuk membuat error dengan format string
// Contoh:
// return fmt.Errorf("umur %d belum cukup", umur)

// ==============================
// 3. ERROR WRAPPING (%w)
// ==============================

// Digunakan untuk membungkus error lain (Go 1.13+)
// Contoh:
// err := errors.New("file tidak ditemukan")
// return fmt.Errorf("gagal membaca file: %w", err)

// ==============================
// 4. errors.Is()
// ==============================

// Digunakan untuk mengecek error tertentu
// Aman untuk error yang sudah di-wrap
// Contoh:
// if errors.Is(err, ErrNotFound) {
//     // handle error
// }

// ==============================
// 5. errors.As()
// ==============================

// Digunakan untuk mengecek & mengambil tipe error tertentu
// Cocok untuk custom error (struct)
// Contoh:
// type MyError struct {
//     Pesan string
// }

// func (e *MyError) Error() string {
//     return e.Pesan
// }

// func main() {
//     err := &MyError{"kesalahan custom"}

//     var myErr *MyError
//     if errors.As(err, &myErr) {
//         fmt.Println("Custom error:", myErr.Pesan)
//     }
// }

// ==============================
// 6. CUSTOM ERROR (STRUCT)
// ==============================

// Custom error harus mengimplementasikan method Error()
// Contoh:
//
// type ValidationError struct {
//     Field string
// }
//
// func (v ValidationError) Error() string {
//     return "field tidak valid"
// }

// ==============================
// 7. POLA FUNGSI ERROR DI GO
// ==============================

// Pola paling umum:
// func namaFungsi() (returnValue, error)

// Contoh:
// data, err := getData()
// if err != nil {
//     return err
// }

// ==============================
// 8. EARLY RETURN (IDIOM GO)
// ==============================

// Error dicek di awal agar kode rapi
// Contoh:
// if err != nil {
//     return err
// }

// ==============================
// 9. PANIC & RECOVER
// ==============================

// panic()
// Digunakan untuk error fatal (program berhenti)
// Tidak disarankan untuk error biasa

// recover()
// Digunakan untuk menangkap panic
// Biasanya dipakai di defer

// ==============================
// 10. RINGKASAN
// ==============================

// errors.New()   -> error statis
// fmt.Errorf()  -> error dengan format
// %w            -> wrap error
// errors.Is()   -> cek error
// errors.As()   -> cek tipe error
// custom error  -> struct + Error()
// panic()       -> error fatal
// recover()     -> tangkap panic
