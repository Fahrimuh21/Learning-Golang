# Belajar Golang Embed

Repositori ini berisi materi pembelajaran tentang fitur `embed` di Go yang memungkinkan menyertakan file statis ke dalam binary saat build.

## Struktur Folder

```
belajar-golang-embed/
├── go.mod                          # File modul Go
├── .gitignore                      # File gitignore
├── main.go                         # Contoh penggunaan embed untuk file, binary, dan directory
├── embed_test.go                   # Test untuk berbagai skenario embed
├── version.txt                     # File teks yang di-embed sebagai string
├── logo.png                        # File binary yang di-embed sebagai []byte
└── files/                          # Folder berisi file teks yang di-embed
    ├── a.txt
    ├── b.txt
    ├── c.txt
    └── d.txt
```

## Deskripsi File

- **go.mod**: Mendefinisikan modul `belajar-golang-embed` dan versi Go.
- **main.go**: Contoh penggunaan `//go:embed` untuk:
  - Meng-embed file teks (`version.txt`) ke dalam variabel `string`.
  - Meng-embed file binary (`logo.png`) ke dalam `[]byte` dan menyimpannya kembali sebagai file baru.
  - Meng-embed seluruh folder (`files/*.txt`) melalui `embed.FS` dan membaca semua file di dalamnya.
- **embed_test.go**: Kumpulan test function yang memperlihatkan:
  - Penggunaan embed string dan byte.
  - Penggunaan embed untuk multiple file secara spesifik dan dengan path glob (`*.txt`).
- **version.txt**: Contoh file teks yang di-embed sebagai string.
- **logo.png**: Contoh file binary yang di-embed sebagai byte slice.
- **files/**: Folder contoh yang berisi beberapa file teks yang di-embed.

## Cara Menjalankan

### Menjalankan Program Utama
```bash
go run main.go
```

Program akan:
- Menampilkan isi `version.txt`.
- Menulis ulang `logo.png` ke file `logo_new.png`.
- Membaca dan menampilkan semua file `.txt` di folder `files`.

### Menjalankan Test
```bash
go test -v
```

Test akan:
- Menampilkan isi `version.txt` dua kali.
- Menulis ulang `logo.png` ke file `logo_new1.png`.
- Membaca isi file `files/a.txt`, `files/b.txt`, `files/c.txt`.
- Membaca semua file `.txt` dalam folder `files`.

## Modul Go

- **Nama Modul**: `belajar-golang-embed`
- **Versi Go**: Disesuaikan dengan `go.mod` (umumnya Go 1.16+ diperlukan untuk fitur embed)

## Persyaratan

- Go 1.16 atau lebih baru (fitur `embed` diperkenalkan di Go 1.16)

## Catatan

- `//go:embed` hanya bekerja di package yang sama dengan deklarasi variabel embed.
- Saat file di-embed, file asli tidak dibutuhkan lagi pada runtime (sudah menjadi bagian dari binary).
- Untuk menulis file ter-embed kembali ke disk, perlu konversi dari `[]byte`.
- Gunakan `embed.FS` untuk meng-embed direktori atau multiple file menggunakan glob (`*.txt`).
