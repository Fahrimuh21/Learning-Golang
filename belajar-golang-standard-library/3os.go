package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}

// ======================================
// PACKAGE os (Operating System) - GOLANG
// ======================================
//
// Package os digunakan untuk berinteraksi langsung
// dengan sistem operasi, seperti:
// - File & directory
// - Permission file
// - Environment variable
// - Argumen command line
// - Proses & exit program
// - Error dari sistem operasi
//
// ======================================
// 1. FILE & DIRECTORY
// ======================================
//
// os.Open("file.txt")
// - Membuka file (read-only)
// - Error jika file tidak ada
//
// os.Create("file.txt")
// - Membuat file baru
// - Menimpa file lama jika sudah ada
//
// os.OpenFile("file.txt", flag, permission)
// Flag umum:
// - os.O_RDONLY : read
// - os.O_WRONLY : write
// - os.O_RDWR   : read & write
// - os.O_CREATE : buat file
// - os.O_APPEND : tambah isi
// - os.O_TRUNC  : hapus isi lama
//
// Permission (Unix style):
// - 0644 : rw-r--r--
// - 0755 : rwxr-xr-x
//
// Selalu tutup file:
// defer file.Close()
//
// ======================================
// 2. OPERASI FILE & FOLDER
// ======================================
//
// os.Remove("file.txt")
// - Menghapus file atau folder kosong
//
// os.RemoveAll("folder")
// - Menghapus folder beserta isinya
//
// os.Rename("lama.txt", "baru.txt")
// - Rename atau pindahkan file
//
// os.Mkdir("folder", 0755)
// - Membuat folder
//
// os.MkdirAll("a/b/c", 0755)
// - Membuat folder bertingkat
//
// os.Stat("file.txt")
// - Mengambil info file
//
// os.IsNotExist(err)
// - Mengecek apakah file tidak ada
//
// ======================================
// 3. ERROR DARI PACKAGE os
// ======================================
//
// Error file system biasanya bertipe:
// *os.PathError
//
// Field penting:
// - Op   : operasi (open, read, write)
// - Path : path file
// - Err  : error aslinya
//
// Gunakan errors.As() untuk mengambil detail error
//
// ======================================
// 4. PERMISSION FILE
// ======================================
//
// os.Chmod("file.txt", 0644)
// - Mengubah permission file
//
// ======================================
// 5. ENVIRONMENT VARIABLE
// ======================================
//
// os.Getenv("KEY")
// - Mengambil nilai environment variable
//
// os.Setenv("KEY", "VALUE")
// - Mengatur environment variable
//
// os.Unsetenv("KEY")
// - Menghapus environment variable
//
// ======================================
// 6. ARGUMEN COMMAND LINE
// ======================================
//
// os.Args
// - os.Args[0] : nama program
// - os.Args[1] : argumen pertama
//
// ======================================
// 7. WORKING DIRECTORY
// ======================================
//
// os.Getwd()
// - Mendapatkan current directory
//
// os.Chdir("path")
// - Mengubah current directory
//
// ======================================
// 8. EXIT PROGRAM
// ======================================
//
// os.Exit(code)
// - 0  : sukses
// - !=0: error
//
// Catatan:
// - defer TIDAK dijalankan saat os.Exit()
//
// ======================================
// 9. STANDARD INPUT / OUTPUT
// ======================================
//
// os.Stdin  : input
// os.Stdout : output
// os.Stderr : error output
//
// Contoh:
// fmt.Fprintln(os.Stderr, "terjadi error")
//
// ======================================
// RINGKASAN
// ======================================
//
// Gunakan package os jika:
// - Berurusan dengan file / folder
// - Mengatur permission
// - Mengambil ENV
// - Membaca argumen CLI
// - Menangani error sistem operasi
//
