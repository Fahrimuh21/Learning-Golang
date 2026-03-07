package belajar_golang_goroutines

//	tentang penggunaan channel pada goroutine
//	dokumentasi channel: https://golang.org/doc/effective_go#channels

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	//tipe data chanel string dan
	// disitu membuat channel baru
	defer close(channel)
	//defer pasti dijalankan pada akhir function
	//jangan lupa di close agar tidak terjadi memory leak

	go func() {
		time.Sleep(2 * time.Second)
		//memberi jeda 2 detik agar untuk memberi waktu channel menerima data
		channel <- "Eko Kurniawan Khannedy"
		// mengirim data ke channel
		//deadlock jika tidak ada yang mengiirim data ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	// menerima data dari channel
	// var data string
	// data <- channel
	//fmt.println("<-channel") sama saja dengan di atas
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	//channel hanya untuk input saja
	//parameter channel bertipe chan<- string
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func OnlyOut(channel <-chan string) {
	//channel hanya untuk output saja
	//parameter channel bertipe <-chan string
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	//menunggu sampai ada data yang di kirim
	//penyimpanan di channel bersifat sementara

	channel := make(chan string, 3)
	//membuat buffered channel dengan kapasitas 3
	//jika gada yang nerima data, maka pengirim data bisa mengirim sesuai kapasitas channel
	//dan tidak akan terjadi deadlock
	defer close(channel)

	//cap(channel) untuk mengetahui kapasitas channel
	//len(channel) untuk mengetahui jumlah data di channel
	fmt.Println("Capacity Channel", cap(channel))
	fmt.Println("Length Channel", len(channel))
	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
		//mengirim 3 data ke channel
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		//menerima 3 data dari channel
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
			//mengirim data ke channel
			//strconv.Itoa untuk mengubah int ke string
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
	//Menerima data Perulangan ke 0
	// Menerima data Perulangan ke 1
	// Menerima data Perulangan ke 2
	// Menerima data Perulangan ke 3
	// Menerima data Perulangan ke 4
	// Menerima data Perulangan ke 5
	// Menerima data Perulangan ke 6
	// Menerima data Perulangan ke 7
	// Menerima data Perulangan ke 8
	// Menerima data Perulangan ke 9
	// Selesai
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		//memilih channel mana yang akan di eksekusi
		case data := <-channel1:
			//jika channel1 yang selesai duluan
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			//jika channel2 yang selesai duluan
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			//biar tidak deadlock
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
			//jika tidak ada channel yang selesai
			//maka akan mengeksekusi default
			time.Sleep(500 * time.Millisecond)
			//beri jeda agar tidak terlalu cepat loopingnya

		}

		if counter == 2 {
			break
		}
	}
}
