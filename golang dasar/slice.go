package main

import "fmt"

//ukuran slice bisa berubah kek vector

func main() {

	//pointer penunjuk data pertama dari array
	//array[low:high]
	//array[:high]
	//array[low:]
	//array[:] => konversi array menjadi slice
	//mulai low, akhir sebelum high

	names := [...]string{"fahri", "dimas", "arif", "elif", "ryan"}
	// slice := names[1:4] // 1 2 3
	// //slice bisa di liat di dokumentasi
	// fmt.Println(slice)

	//DEKLARASI MANUAL
	var slice1 []string = names[:]

	//FUNCTION
	//1. len (panjang)
	fmt.Println(len(slice1))
	//2. cap (kapasitas)
	fmt.Println(cap(slice1))

	//3. append(slice, data) menambahkan data ke posisi terakhir
	//jika cap sudah penuh maka akan membuat array yang baru
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "MInggu baru"
	fmt.Println(days) //ikut berubah

	//kalo di append
	dayslice2 := append(daySlice1, "libur baru")
	dayslice2[0] = "ups"

	fmt.Println(dayslice2) //ups, sabtu baru, libur baru
	fmt.Println(days)      //ga ikut berubah
	fmt.Println(daySlice1)

	//4. make([]TYPE DATA, length, capacity) membuat slice baru
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "fahri"
	newSlice[1] = "waud"
	// newSlice[2] = "wai" tidak bisa, harus append karna len=2 harus pake append

	fmt.Println(newSlice)

	//5. copy(destination, source) menyalin slice dari source ke destination
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

}
