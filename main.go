package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", mainRouter)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
