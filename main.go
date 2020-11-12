package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr, path string
	flag.StringVar(&addr, "addr", ":8080", "address to run server on")
	flag.StringVar(&path, "path", ".", "directory to serve")
	flag.Parse()

	log.Printf("Serving ./ on %s\n", addr)
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
