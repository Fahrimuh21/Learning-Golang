# Belajar Golang Context

Repositori ini berisi materi pembelajaran tentang Context package dalam bahasa pemrograman Go. Context digunakan untuk mengatur timeout, cancellation, dan komunikasi data antar goroutine secara aman.

## Struktur Folder

```
belajar-golang-context/
├── go.mod                          # File modul Go
├── .gitignore                      # File gitignore
└── context_test.go                 # Test file untuk berbagai konsep context
```

## Deskripsi File

- **go.mod**: Mendefinisikan modul `belajar-golang-context` dengan Go versi 1.15.
- **context_test.go**: Berisi berbagai test function yang mendemonstrasikan penggunaan context:
  - `TestContext`: Basic context (Background dan TODO)
  - `TestContextWithValue`: Context dengan value untuk passing data
  - `TestContextWithCancel`: Context dengan cancellation manual
  - `TestContextWithTimeout`: Context dengan timeout duration
  - `TestContextWithDeadline`: Context dengan deadline waktu tertentu
  - `CreateCounter`: Helper function yang menggunakan context untuk cancel goroutine

## Cara Menjalankan Test

### Menjalankan Semua Test
```bash
go test -v
```

### Menjalankan Test Spesifik
```bash
go test -v -run=TestContext
```

### Menjalankan dengan Timeout
```bash
go test -v -timeout=30s
```

## Modul Go

- **Nama Modul**: `belajar-golang-context`
- **Versi Go**: 1.15

## Persyaratan

- Go versi 1.15 atau lebih baru

## Konsep yang Dipelajari

1. **Context Background**: Context root yang kosong, digunakan sebagai parent context
2. **Context TODO**: Context kosong untuk situasi ketika belum tahu context yang akan digunakan
3. **Context WithValue**: Menambahkan key-value pair ke context untuk passing data antar goroutine
4. **Context WithCancel**: Membuat context yang dapat dibatalkan secara manual
5. **Context WithTimeout**: Membuat context dengan batas waktu duration
6. **Context WithDeadline**: Membuat context dengan batas waktu absolut
7. **Cancellation Propagation**: Ketika parent context dibatalkan, semua child context juga dibatalkan
8. **Goroutine Cancellation**: Menggunakan context untuk menghentikan goroutine dengan aman

## Catatan

Context sangat penting dalam Go untuk:
- Mengatur timeout pada operasi I/O
- Membatalkan operasi ketika tidak lagi diperlukan
- Passing request-scoped data antar goroutine
- Mengelola lifecycle goroutine dalam aplikasi concurrent

File test ini mendemonstrasikan berbagai cara penggunaan context dalam skenario praktis. Beberapa test menggunakan `time.Sleep()` sehingga mungkin memerlukan waktu eksekusi yang lebih lama.