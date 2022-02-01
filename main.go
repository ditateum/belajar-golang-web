package main

import (
	"belajar-golang-web/handler"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler.Handler,
	}

	fmt.Println("Server running")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
