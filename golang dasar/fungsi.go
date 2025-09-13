package main

import "fmt"

//DEKLARASI

//single return values 
func tambah(a int, b int) int {//gunanya untuk dimainkan di fungsi utama 
	return a + b
}

//multiple return values
func divmod(a, b int) (int, int, int) {
    return a / b, a % b, a+b
}

//named return values
func luasPersegi(s int) (luas int) {
    luas = s * s
    return // return luas
}
//short declaration
func add(a, b int) int { return a + b }

//bagian dari menghiraukan fungsi 
func getData() (int, string) {
    return 100, "Hello"
}

func main() {
	fmt.Println(tambah(2, 3))
	fmt.Println(divmod(4, 3))

	//menghiraukan fungsi 
	x, y := getData()
    fmt.Println(x, y)  // pakai keduanya

    // Kalau cuma butuh string-nya
    _, msg := getData()
    fmt.Println(msg)   // Output: Hello



}
