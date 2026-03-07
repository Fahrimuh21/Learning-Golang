package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
}

// ======================================
// PACKAGE strconv - GOLANG
// ======================================
//
// Package strconv digunakan untuk:
// - Konversi STRING <-> TIPE DATA LAIN
// - Parsing angka dari string
// - Mengubah angka menjadi string
// - Mengubah boolean ke string dan sebaliknya
//
// Umum dipakai saat:
// - Input dari user (CLI / ENV / HTTP)
// - Membaca file / JSON manual
// - Parsing os.Args atau flag
//
// ======================================
// 1. STRING -> INTEGER
// ======================================
//
// strconv.Atoi(s string) (int, error)
// - Mengubah string ke int
// - Error jika string bukan angka
//
// Contoh:
// i, err := strconv.Atoi("123")
//
// ======================================
// 2. INTEGER -> STRING
// ======================================
//
// strconv.Itoa(i int) string
// - Mengubah int ke string
//
// Contoh:
// s := strconv.Itoa(123)
//
// ======================================
// 3. STRING -> ANGKA (FUNGSI UMUM)
// ======================================
//
// strconv.ParseInt(s, base, bitSize)
// - base    : basis bilangan (2, 8, 10, 16)
// - bitSize : 8, 16, 32, 64
//
// Contoh:
// n, err := strconv.ParseInt("1010", 2, 64)
//
// strconv.ParseFloat(s, bitSize)
// - bitSize : 32 atau 64
//
// Contoh:
// f, err := strconv.ParseFloat("3.14", 64)
//
// ======================================
// 4. ANGKA -> STRING (FORMAT)
// ======================================
//
// strconv.FormatInt(i, base)
// - base: 2 (biner), 8 (oktal), 10, 16 (hex)
//
// Contoh:
// s := strconv.FormatInt(10, 2) // "1010"
//
// strconv.FormatFloat(f, fmt, prec, bitSize)
// - fmt  : 'f', 'e', 'g'
// - prec : jumlah digit
//
// Contoh:
// s := strconv.FormatFloat(3.14, 'f', 2, 64)
//
// ======================================
// 5. STRING -> BOOLEAN
// ======================================
//
// strconv.ParseBool(s string) (bool, error)
//
// String yang dianggap TRUE:
// "1", "t", "T", "TRUE", "true", "True"
//
// String yang dianggap FALSE:
// "0", "f", "F", "FALSE", "false", "False"
//
// ======================================
// 6. BOOLEAN -> STRING
// ======================================
//
// strconv.FormatBool(b bool) string
//
// ======================================
// 7. QUOTE & UNQUOTE STRING
// ======================================
//
// strconv.Quote(s string)
// - Menambahkan tanda kutip & escape karakter
//
// strconv.Unquote(s string)
// - Menghapus tanda kutip
//
// ======================================
// 8. ERROR HANDLING
// ======================================
//
// Semua fungsi parsing mengembalikan error
// WAJIB dicek
//
// Contoh:
// i, err := strconv.Atoi(input)
// if err != nil {
//     // input bukan angka
// }
//
// ======================================
// 9. PERBANDINGAN FUNGSI UMUM
// ======================================
//
// Atoi  <-> Itoa
// Parse <-> Format
//
// Atoi / Itoa -> khusus int
// Parse / Format -> fleksibel (base & bit)
//
// ======================================
// 10. RINGKASAN CEPAT
// ======================================
//
// strconv.Atoi()       -> string ke int
// strconv.Itoa()       -> int ke string
// strconv.ParseInt()  -> string ke angka
// strconv.FormatInt() -> angka ke string
// strconv.ParseFloat()-> string ke float
// strconv.ParseBool() -> string ke bool
//
// Gunakan strconv untuk:
// - Parsing input user
// - Konversi tipe data
// - Menghindari manual casting
//
