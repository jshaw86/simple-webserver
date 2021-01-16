package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080"}

    http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {

        w.Write([]byte("ok"))
    })

	log.Printf("Listening https://0.0.0.0:8080")
	log.Fatal(srv.ListenAndServe())
}


