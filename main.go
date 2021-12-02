package main

import (
	"fmt"
	"gow"
	"net/http"
)

func main() {
	r := gow.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	r.Run(":9000")
}
