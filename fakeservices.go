package main

import (
	"fmt"
	"net/http"
)

func standard_response ( response http.ResponseWriter, request *http.Request ) {
	fmt.Fprintf(response, "hello and have a nice day :-)\n")
}

func show_headers ( response http.ResponseWriter, request *http.Request ) {
	for name, content := range request.Header {
		fmt.Fprintf(response, "%-30s: %s\n", name, content)
	}
}


func main () {
	http.HandleFunc( "/", standard_response )
	http.HandleFunc( "/header", show_headers )
	http.ListenAndServe(":4000", nil)
}