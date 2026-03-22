package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	//untuk memanggil method yang ada di struct, kita bisa menggunakan dot (.) lalu nama methodnya, lalu kita bisa mengirim parameter ke method tersebut dengan menggunakan tanda kurung dan isi parameter yang ingin kita kirim, seperti "Budi" pada contoh di atas.
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Eko",
	})
	//jadi pada contoh di atas, kita membuat template dengan nama "FUNCTION" yang berisi pemanggilan method SayHello dengan parameter "Budi", lalu kita kirim data ke template tersebut dengan menggunakan struct MyPage yang memiliki field Name dengan nilai "Eko". Sehingga ketika template dieksekusi, maka akan menghasilkan output "Hello Budi, My Name Is Eko".
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	//untuk membuat function global, kita bisa menggunakan method Funcs pada template, lalu kita bisa mengirimkan map[string]interface{} yang berisi nama function dan implementasi function tersebut, seperti contoh di bawah ini.
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Eko",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	//register function global dengan nama "upper" yang akan mengubah string menjadi huruf kapital, lalu kita bisa menggunakan function tersebut di dalam template dengan cara memanggil nama function tersebut, seperti contoh di bawah ini.
	t = template.Must(t.Parse(`{{ upper .Name }}`))
	//jadi pada contoh di atas, kita membuat function global dengan nama "upper" yang akan mengubah string menjadi huruf kapital, lalu kita parse template yang berisi pemanggilan function "upper" dengan parameter .Name, lalu kita kirim data ke template tersebut dengan menggunakan struct MyPage yang memiliki field Name dengan nilai "Eko Kurniawan Khannedy". Sehingga ketika template dieksekusi, maka akan menghasilkan output "EKO KURNIAWAN KHANNEDY".

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Eko Kurniawan Khannedy",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	//hasil sayhello akan diubah menjadi huruf kapital dengan menggunakan pipeline, sehingga output yang dihasilkan adalah "HELLO EKO KURNIAWAN KHANNEDY".

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Eko Kurniawan Khannedy",
	})
}

func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
