# Belajar Golang Standard Library

Repositori ini berisi materi pembelajaran tentang Standard Library bahasa pemrograman Go (Golang). Folder ini mencakup berbagai package bawaan Go yang penting untuk pengembangan aplikasi, mulai dari formatting, error handling, hingga operasi file dan data.

## Struktur Folder

```
belajar-golang-standard-library/
├── go.mod                          # File modul Go
├── .gitignore                      # File gitignore
├── 1fmt.go                         # Package fmt untuk formatting
├── 2errors.go                      # Package errors untuk error handling
├── 3os.go                          # Package os untuk operasi sistem
├── 4flag.go                        # Package flag untuk command line arguments
├── 5strings.go                     # Package strings untuk manipulasi string
├── 6strconv.go                     # Package strconv untuk konversi string
├── 7math.go                        # Package math untuk operasi matematika
├── 8list.go                        # Package container/list untuk linked list
├── 9ring.go                        # Package container/ring untuk circular list
├── 10sort.go                       # Package sort untuk sorting data
├── 11time.go                       # Package time untuk operasi waktu
├── 12duration.go                   # Package time untuk duration
├── 13reflect.go                    # Package reflect untuk reflection
├── 14regexp.go                     # Package regexp untuk regular expression
├── 15base64.go                     # Package encoding/base64 untuk base64 encoding
├── 16csv_reader.go                 # Package encoding/csv untuk membaca CSV
├── 17csv_writer.go                 # Package encoding/csv untuk menulis CSV
├── 18slices.go                     # Package slices untuk operasi slice
├── 19path.go                       # Package path/filepath untuk path manipulation
├── 20filepath.go                   # Package path/filepath untuk file path
├── 21bufio_read.go                 # Package bufio untuk buffered reading
├── 22bufio_write.go                # Package bufio untuk buffered writing
└── 23file.go                       # Operasi file dengan os dan bufio
```

## Deskripsi File

Berikut adalah deskripsi singkat untuk setiap file dalam folder ini:

1. **1fmt.go**: Penggunaan package fmt untuk formatting output dan input.
2. **2errors.go**: Membuat dan menggunakan error dengan package errors.
3. **3os.go**: Operasi sistem seperti mendapatkan arguments command line dan hostname menggunakan package os.
4. **4flag.go**: Parsing command line flags menggunakan package flag.
5. **5strings.go**: Manipulasi string seperti contains, split, toLower, toUpper, trim, dan replace menggunakan package strings.
6. **6strconv.go**: Konversi antara string dan tipe data lain menggunakan package strconv.
7. **7math.go**: Operasi matematika seperti ceil, floor, round, max, min menggunakan package math.
8. **8list.go**: Penggunaan linked list dengan package container/list.
9. **9ring.go**: Penggunaan circular list (ring) dengan package container/ring.
10. **10sort.go**: Sorting data dengan package sort, termasuk custom sorting untuk struct.
11. **11time.go**: Operasi waktu seperti mendapatkan waktu sekarang dan parsing/formatting waktu menggunakan package time.
12. **12duration.go**: Penggunaan duration dalam package time untuk menghitung selisih waktu.
13. **13reflect.go**: Reflection untuk menganalisis tipe data dan struct pada runtime menggunakan package reflect.
14. **14regexp.go**: Regular expression untuk pencocokan pola string menggunakan package regexp.
15. **15base64.go**: Encoding dan decoding base64 menggunakan package encoding/base64.
16. **16csv_reader.go**: Membaca data CSV menggunakan package encoding/csv.
17. **17csv_writer.go**: Menulis data ke format CSV menggunakan package encoding/csv.
18. **18slices.go**: Operasi pada slice seperti min, max, contains, index menggunakan package slices.
19. **19path.go**: Manipulasi path file menggunakan package path/filepath.
20. **20filepath.go**: Operasi path file seperti mendapatkan directory, base name, extension menggunakan package path/filepath.
21. **21bufio_read.go**: Membaca data secara buffered per line menggunakan package bufio.
22. **22bufio_write.go**: Menulis data secara buffered menggunakan package bufio.
23. **23file.go**: Operasi file seperti create, read, write, append menggunakan package os dan bufio.

## Cara Menjalankan

Untuk menjalankan salah satu file, gunakan perintah berikut di terminal:

```bash
go run nama_file.go
```

Contoh:

```bash
go run 1fmt.go
```

Beberapa file mungkin memerlukan input atau argumen command line, seperti 4flag.go yang menggunakan flag parsing.

## Modul Go

File `go.mod` mendefinisikan modul ini dengan nama `belajar-golang-standard-library` dan menggunakan Go versi 1.21.3.

## Persyaratan

- Go versi 1.21.3 atau lebih baru
- Terminal atau command prompt untuk menjalankan program

## Catatan

File-file ini dirancang untuk mempelajari berbagai package standard library Go yang sering digunakan dalam pengembangan aplikasi. Setiap file fokus pada satu package atau konsep tertentu untuk memudahkan pemahaman. Beberapa file mungkin memerlukan file tambahan atau input khusus untuk demonstrasi lengkap.