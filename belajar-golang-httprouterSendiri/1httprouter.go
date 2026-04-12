package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})
	// request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	// recorder := httptest.NewRecorder()

	// router.ServeHTTP(recorder, request)
	// response := recorder.Result()

	// bytes, _ := io.ReadAll(response.Body)

	// assert.Equal(t,"Hello Get", string(bytes))
	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
