# Belajar Golang Goroutines

Repositori ini berisi materi pembelajaran tentang Goroutines dan Concurrency dalam bahasa pemrograman Go. Folder ini mencakup berbagai konsep concurrency mulai dari basic goroutine, channel, hingga synchronization primitives.

## Struktur Folder

```
belajar-golang-goroutines/
├── go.mod                          # File modul Go
├── 1goroutine_test.go              # Basic goroutine creation
├── 2channel_test.go                # Channel untuk komunikasi antar goroutine
├── 3race_condition_test.go         # Race condition problem
├── 4mutex_test.go                  # Mutex untuk synchronization
├── 5waitgroup_test.go              # WaitGroup untuk menunggu goroutine
├── 6once_test.go                   # sync.Once untuk eksekusi sekali
├── 7pool_test.go                   # sync.Pool untuk object pooling
├── 8map_test.go                    # sync.Map untuk concurrent map
├── 9cond_test.go                   # sync.Cond untuk conditional synchronization
├── 10atomic_test.go                # Atomic operations
├── 11timer_test.go                 # Timer untuk scheduling
├── 12ticker_test.go                # Ticker untuk periodic tasks
└── 13gomaxprocs_test.go            # GOMAXPROCS untuk thread management
```

## Deskripsi File

1. **1goroutine_test.go**: Demonstrasi dasar pembuatan dan eksekusi goroutine, termasuk menjalankan banyak goroutine secara bersamaan.
2. **2channel_test.go**: Penggunaan channel sebagai medium komunikasi antar goroutine, termasuk buffered dan unbuffered channel.
3. **3race_condition_test.go**: Menunjukkan masalah race condition ketika multiple goroutine mengakses shared variable tanpa synchronization.
4. **4mutex_test.go**: Penggunaan mutex (sync.Mutex dan sync.RWMutex) untuk mengatasi race condition dengan locking mechanism.
5. **5waitgroup_test.go**: Menggunakan sync.WaitGroup untuk menunggu semua goroutine selesai sebelum program utama berakhir.
6. **6once_test.go**: sync.Once untuk memastikan fungsi hanya dieksekusi sekali meskipun dipanggil oleh banyak goroutine.
7. **7pool_test.go**: sync.Pool untuk reuse object dan mengurangi alokasi memory berulang.
8. **8map_test.go**: sync.Map untuk map yang thread-safe tanpa perlu external locking.
9. **9cond_test.go**: sync.Cond untuk conditional synchronization, menunggu kondisi tertentu sebelum melanjutkan.
10. **10atomic_test.go**: Operasi atomic untuk primitive types tanpa locking, lebih efisien untuk operasi sederhana.
11. **11timer_test.go**: Penggunaan time.Timer, time.After, dan time.AfterFunc untuk scheduling task di masa depan.
12. **12ticker_test.go**: time.Ticker dan time.Tick untuk menjalankan task secara periodik/berulang.
13. **13gomaxprocs_test.go**: Mengatur jumlah maksimum thread yang digunakan Go runtime dengan runtime.GOMAXPROCS.

## Cara Menjalankan Test

### Menjalankan Semua Test
```bash
go test -v
```

### Menjalankan Test Spesifik
```bash
go test -v -run=TestCreateGoroutine
```

### Menjalankan Test dengan Race Detection
```bash
go test -v -race
```

### Menjalankan Test dengan Timeout
```bash
go test -v -timeout=30s
```

## Modul Go

- **Nama Modul**: `belajar-golang-goroutines`
- **Versi Go**: 1.15

## Persyaratan

- Go versi 1.15 atau lebih baru

## Konsep yang Dipelajari

1. **Goroutine**: Lightweight thread untuk concurrent execution
2. **Channel**: Typed conduit untuk communication between goroutines
3. **Race Condition**: Problem ketika multiple goroutines mengakses shared data
4. **Synchronization Primitives**:
   - Mutex (Mutual Exclusion)
   - WaitGroup
   - Once
   - Cond
5. **Concurrent Data Structures**:
   - sync.Map
   - sync.Pool
6. **Atomic Operations**: Lock-free operations untuk primitive types
7. **Timer & Ticker**: Scheduling dan periodic tasks
8. **Runtime Control**: Mengatur GOMAXPROCS untuk thread management

## Catatan

File-file ini menggunakan testing framework Go untuk mendemonstrasikan konsep concurrency. Beberapa test mungkin memerlukan waktu eksekusi yang lebih lama karena menggunakan `time.Sleep()`. Gunakan flag `-race` saat menjalankan test untuk mendeteksi race condition. File-file ini dirancang untuk dipelajari secara berurutan dari basic hingga advanced concepts.