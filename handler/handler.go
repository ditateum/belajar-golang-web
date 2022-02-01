package handler

import (
	"fmt"
	"net/http"
)

var Handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world")
}
