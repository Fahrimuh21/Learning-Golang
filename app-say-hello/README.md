# App Say Hello

Repositori ini berisi contoh aplikasi Go yang menggunakan modul eksternal `go-say-hello` untuk pembelajaran tentang Go Modules dan dependency management.

## Struktur Folder

```
app-say-hello/
├── go.mod                          # File modul Go dengan dependency
└── main.go                         # Aplikasi utama yang menggunakan modul go-say-hello
```

## Deskripsi File

- **go.mod**: Mendefinisikan modul `github.com/ProgrammerZamanNow/app-say-hello` dengan dependency `github.com/ProgrammerZamanNow/go-say-hello/v2 v2.0.0`.
- **main.go**: Aplikasi yang mengimport dan menggunakan fungsi `SayHello` dari modul `go-say-hello`.

## Cara Menjalankan

1. Pastikan dependency terinstall:
   ```bash
   go mod tidy
   ```

2. Jalankan aplikasi:
   ```bash
   go run main.go
   ```

Output yang diharapkan:
```
Hello Eko
```

## Modul Go

- **Nama Modul**: `github.com/ProgrammerZamanNow/app-say-hello`
- **Versi Go**: 1.15
- **Dependency**: `github.com/ProgrammerZamanNow/go-say-hello/v2 v2.0.0`

## Cara Upgrade Modul

Untuk upgrade versi modul dependency:

```bash
go get github.com/ProgrammerZamanNow/go-say-hello/v2@latest
```

Atau versi spesifik:
```bash
go get github.com/ProgrammerZamanNow/go-say-hello/v2@v2.1.0
```

## Persyaratan

- Go versi 1.15 atau lebih baru
- Akses internet untuk download dependency (jika belum ada di cache)

## Catatan

Ini adalah contoh aplikasi yang menggunakan modul eksternal untuk mempelajari konsep dependency management dalam Go Modules. Aplikasi ini mendemonstrasikan cara mengimport dan menggunakan package dari modul lain.