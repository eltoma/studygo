package main

import (
	"net/http"
	"fmt"
)

/**
 * http server
 */
func main() {
	// 指针类型
	http.HandleFunc("/", func(writer http.ResponseWriter,
		request *http.Request) {

		fmt.Fprintf(writer, "<h1>Hello world %s!</h1>",
			request.FormValue("name"))
	})

	http.ListenAndServe(":8888", nil)
}
