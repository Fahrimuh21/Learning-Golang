package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` //tag untuk json atau validasi
}

type Person struct {
	Name    string `required:"true" max:"10"` //required untuk data ada dan tidak ada lebih dari 10 karakter
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) { //baca struct field
	var valueType reflect.Type = reflect.TypeOf(value) //type nya Type dari reflect
	fmt.Println("Type Name", valueType.Name())         //nama typenya contoh sini Sample atau Person
	for i := 0; i < valueType.NumField(); i++ {        //jumlah elemen field daari struct nya
		var structField reflect.StructField = valueType.Field(i)     //dapatkan field nya satu2
		fmt.Println(structField.Name, "with type", structField.Type) //nama field dan type data field contohnnya structField.name = Name, structField.Type = string
		//tag
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Eko"})         //memanggil function readField dengan mengirimkan struct Sample
	readField(Person{"Eko", "", ""}) //memanggil function readField dengan mengirimkan struct Person

	person := Person{
		Name:    "ada",
		Address: "",
		Email:   "ada",
	}
	fmt.Println(IsValid(person))
}
