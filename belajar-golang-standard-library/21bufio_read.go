package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n") // 2 line string
	reader := bufio.NewReader(input)                                   //baca dari input

	for {
		line, _, err := reader.ReadLine() //baca per line
		if err == io.EOF {                //jika sudah akhir file
			break
		}

		fmt.Println(string(line)) //tampilkan line
	}
}
