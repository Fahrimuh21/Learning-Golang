package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir(filepath.FromSlash("hello/world.go")))
	fmt.Println(filepath.Base(filepath.FromSlash("hello/world.go")))
	fmt.Println(filepath.Ext(filepath.FromSlash("hello/world.go")))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
