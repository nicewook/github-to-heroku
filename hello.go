package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
		log.Println("$PORT is not set, so port is set to ", port)
		log.Printf("Ctrl+Click http://localhost:8888")
	}

	if err := http.ListenAndServe(":"+port, http.HandlerFunc(helloHandler)); err != nil {
		log.Fatal(err)
	}
}
