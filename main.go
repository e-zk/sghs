package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr, path string
	flag.StringVar(&addr, "a", ":8080", "address:port to serve from")
	flag.StringVar(&path, "p", ".", "directory to serve")
	flag.Parse()

	log.Printf("Serving %s on %s\n", path, addr)
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
