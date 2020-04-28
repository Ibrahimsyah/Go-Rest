package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":5000", nil)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
