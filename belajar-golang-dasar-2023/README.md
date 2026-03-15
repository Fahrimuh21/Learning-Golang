# Belajar Golang Dasar 2023

Repositori ini berisi materi pembelajaran dasar bahasa pemrograman Go (Golang) untuk tahun 2023. Folder ini mencakup berbagai konsep fundamental dari Go, mulai dari hello world hingga topik lanjutan seperti pointer, error handling, dan struktur data.

## Struktur Folder

```
belajar-golang-dasar-2023/
├── go.mod                          # File modul Go
├── .gitignore                      # File gitignore
├── sample.go                       # File contoh sederhana
├── 1helloworld.go                  # Program Hello World
├── 2number.go                      # Penggunaan tipe data number
├── 3boolean.go                     # Penggunaan tipe data boolean
├── 4string.go                      # Penggunaan tipe data string
├── 5variable.go                    # Deklarasi dan penggunaan variabel
├── 6constant.go                    # Penggunaan konstanta
├── 7conversion.go                  # Konversi tipe data
├── 8type_declaration.go            # Deklarasi tipe data kustom
├── 9matematika.go                  # Operasi matematika
├── 10perbandingan.go               # Operasi perbandingan
├── 11operasi_boolean.go            # Operasi boolean (AND, OR, NOT)
├── 12array.go                      # Penggunaan array
├── 13slice.go                      # Penggunaan slice
├── 14map.go                        # Penggunaan map
├── 15if.go                         # Statement if-else
├── 16switch.go                     # Statement switch
├── 17for.go                        # Loop for
├── 18break.go                      # Statement break
├── 19continue.go                   # Statement continue
├── 20function.go                   # Fungsi dasar
├── 21function_parameter.go         # Fungsi dengan parameter
├── 22function_return_value.go      # Fungsi dengan return value
├── 23return_multiple_values.go     # Fungsi dengan multiple return values
├── 24named_return_values.go        # Named return values
├── 25variadic_function.go          # Variadic function
├── 26function_value.go             # Function as value
├── 27function_as_parameter.go      # Function as parameter
├── 28anonymous_function.go         # Anonymous function
├── 29recursive_function.go         # Recursive function
├── 30closure.go                    # Closure
├── 31defer.go                      # Defer statement
├── 32panic.go                      # Panic
├── 33struct.go                     # Struct
├── 34interface.go                  # Interface
├── 35comment.go                    # Komentar
├── 36any.go                        # Tipe data any (interface{})
├── 37nil.go                        # Nil
├── 38type_assertions.go            # Type assertions
├── 39pointer.go                    # Pointer
├── 40asterisk_operator.go          # Asterisk operator
├── 41new.go                        # Fungsi new()
├── 42pointer_function.go           # Pointer di function
├── 43pointer_method.go             # Pointer di method
├── 44import.go                     # Import package
├── 45init.go                       # Init function
├── 46error.go                      # Error handling
├── 47error_custom.go               # Custom error
├── database/                       # Package database
│   └── mysql.go                    # Koneksi database MySQL
├── helper/                         # Package helper
│   └── helper.go                   # Fungsi helper
└── internal/                       # Package internal
    └── internal.go                 # Fungsi internal
```

## Deskripsi File

Berikut adalah deskripsi singkat untuk setiap file dalam folder ini:

1. **1helloworld.go**: Program sederhana yang mencetak "Hello World" ke konsol.
2. **2number.go**: Demonstrasi penggunaan tipe data numerik (int, float).
3. **3boolean.go**: Penggunaan tipe data boolean (true/false).
4. **4string.go**: Operasi pada string, termasuk panjang dan akses karakter.
5. **5variable.go**: Deklarasi variabel dengan berbagai cara.
6. **6constant.go**: Penggunaan konstanta dan sifatnya yang tidak dapat diubah.
7. **7conversion.go**: Konversi antara tipe data yang berbeda.
8. **8type_declaration.go**: Membuat tipe data kustom menggunakan type declaration.
9. **9matematika.go**: Operasi matematika dasar dan augmented assignment.
10. **10perbandingan.go**: Operasi perbandingan (==, !=, <, >, dll).
11. **11operasi_boolean.go**: Operasi logika boolean (&&, ||, !).
12. **12array.go**: Penggunaan array dengan ukuran tetap.
13. **13slice.go**: Penggunaan slice sebagai array dinamis.
14. **14map.go**: Penggunaan map untuk menyimpan data key-value.
15. **15if.go**: Kontrol alur dengan statement if-else.
16. **16switch.go**: Kontrol alur dengan statement switch-case.
17. **17for.go**: Loop menggunakan for dengan berbagai bentuk.
18. **18break.go**: Menggunakan break untuk menghentikan loop.
19. **19continue.go**: Menggunakan continue untuk melanjutkan iterasi loop.
20. **20function.go**: Definisi dan pemanggilan fungsi dasar.
21. **21function_parameter.go**: Fungsi dengan parameter input.
22. **22function_return_value.go**: Fungsi yang mengembalikan nilai.
23. **23return_multiple_values.go**: Fungsi yang mengembalikan multiple values.
24. **24named_return_values.go**: Menggunakan named return values.
25. **25variadic_function.go**: Fungsi dengan parameter variadic.
26. **26function_value.go**: Menggunakan fungsi sebagai value.
27. **27function_as_parameter.go**: Melewatkan fungsi sebagai parameter.
28. **28anonymous_function.go**: Fungsi anonim (anonymous function).
29. **29recursive_function.go**: Fungsi rekursif.
30. **30closure.go**: Konsep closure dalam Go.
31. **31defer.go**: Menggunakan defer untuk menunda eksekusi.
32. **32panic.go**: Menyebabkan panic dalam program.
33. **33struct.go**: Definisi dan penggunaan struct.
34. **34interface.go**: Penggunaan interface.
35. **35comment.go**: Cara menggunakan komentar dalam kode.
36. **36any.go**: Menggunakan interface{} sebagai any type.
37. **37nil.go**: Konsep nil dalam Go.
38. **38type_assertions.go**: Type assertions untuk interface.
39. **39pointer.go**: Penggunaan pointer.
40. **40asterisk_operator.go**: Operator asterisk (*) untuk pointer.
41. **41new.go**: Fungsi new() untuk alokasi memory.
42. **42pointer_function.go**: Pointer dalam parameter fungsi.
43. **43pointer_method.go**: Pointer dalam method struct.
44. **44import.go**: Cara mengimport package.
45. **45init.go**: Fungsi init() yang dijalankan otomatis.
46. **46error.go**: Error handling dasar.
47. **47error_custom.go**: Membuat custom error.

### Package Tambahan

- **database/mysql.go**: Package untuk koneksi database MySQL dengan init function.
- **helper/helper.go**: Package helper dengan fungsi publik dan privat.
- **internal/internal.go**: Package internal dengan init function.

## Cara Menjalankan

Untuk menjalankan salah satu file, gunakan perintah berikut di terminal:

```bash
go run nama_file.go
```

Contoh:

```bash
go run 1helloworld.go
```

Untuk file dalam package lain, pastikan untuk menjalankan dari root folder dan menggunakan import yang sesuai.

## Modul Go

File `go.mod` mendefinisikan modul ini dengan nama `belajar-golang-dasar` dan menggunakan Go versi 1.21.3.

## Persyaratan

- Go versi 1.21.3 atau lebih baru
- Terminal atau command prompt untuk menjalankan program

## Catatan

File-file ini dirancang untuk pembelajaran bertahap, dimulai dari konsep paling dasar hingga topik yang lebih kompleks. Setiap file dapat dijalankan secara independen untuk memahami konsep tertentu.