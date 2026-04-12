package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Gak Ketemu")
	})
	// //pakai interface bukan params
	// router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// 	panic("ups")
	// })

	request := httptest.NewRequest("GET", "http://localhost:3000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	Body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(Body))
}
