package main

import "net/http"

func main() {
	http.HandleFunc("/", helloGoHandler)
}

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}

