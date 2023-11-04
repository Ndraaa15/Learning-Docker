package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Docker!"))
	})

	http.ListenAndServe(port, nil)
	log.Println("Listening on port", port)
}
