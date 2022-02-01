package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func ServerTest(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server run on port 8080")
}
