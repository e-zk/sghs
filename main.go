package main

import (
	"flag"
	"net/http"
)

func main() {
	var addr, path string
	flag.StringVar(&addr, "addr", ":8080", "address to run server on")
	flag.StringVar(&path, "path", ".", "directory to serve")
	flag.Parse()

	print("Serving ./ on " + addr + "\n")
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
