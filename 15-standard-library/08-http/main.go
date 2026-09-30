package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, ", r.URL.Query().Get("name"))
	})

	request := httptest.NewRequest(http.MethodGet, "/?name=Go", nil)
	response := httptest.NewRecorder() // Capture the reply without a server.
	handler.ServeHTTP(response, request)

	fmt.Println("status:", response.Code)
	fmt.Println("body:", response.Body.String())
}
