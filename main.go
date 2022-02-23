package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

func main() {
	var addr, path, tls, certFile, keyFile string
	flag.StringVar(&addr, "l", ":8080", "[address]:port to listen on")
	flag.StringVar(&path, "p", ".", "path to serve")
	flag.StringVar(&tls, "t", "", "tls cert_path:key_path")
	flag.Parse()

	if tls != "" {
		s := strings.Split(tls, ":")
		if len(s) < 2 || s[0] == "" || s[1] == "" {
			log.Fatal("Invalid cert/key path.")
		}

		// TODO check if files exist
		certFile = s[0]
		keyFile = s[1]
	}

	// TODO custom handler that logs requests
	// TODO -q (quiet) flag (no logging)
	http.Handle("/", http.FileServer(http.Dir("./")))

	if tls != "" {
		log.Printf("Serving %s on %s (tls=%s:%s).\n", path, addr, certFile, keyFile)
		if err := http.ListenAndServeTLS(addr, certFile, keyFile, nil); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Serving %s on %s.\n", path, addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}
}
