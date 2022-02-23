package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	addr, path, tls string
	quiet           bool
)

// "middleware" that logs the request details then passes it on
func LogHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// print to stdout
		if !quiet {
			log.Printf("[%s]\t%s\t%s\n", r.RemoteAddr, r.Method, r.URL.Path)
		}
		h.ServeHTTP(w, r)
	})
}

func init() {
	flag.StringVar(&addr, "l", ":8080", "[address]:port to listen on")
	flag.BoolVar(&quiet, "q", false, "quiet mode (don't print logs)")
	flag.StringVar(&path, "p", ".", "path to serve")
	flag.StringVar(&tls, "t", "", "tls cert and key path in the format of: 'cert_path:key_path'")
	flag.Parse()
}

func main() {
	// if tls cert has been given, we need to check if the files exist
	// and split the given string into two variables:
	var certFile, keyFile string
	if tls != "" {
		s := strings.Split(tls, ":")
		if len(s) < 2 || s[0] == "" || s[1] == "" {
			log.Fatal("Invalid cert/key path")
		}

		certFile = s[0]
		_, err := os.Open(certFile)
		if os.IsNotExist(err) {
			log.Fatal("Given cert file does not exist")
		}
		keyFile = s[1]
		_, err = os.Open(keyFile)
		if os.IsNotExist(err) {
			log.Fatal("Given cert key file does not exist")
		}
	}

	// use the LogHandler
	http.Handle("/", LogHandler(http.FileServer(http.Dir(path))))

	// if tls cert has been given, serve tls
	// otherwise serve plain http
	if tls != "" {
		log.Printf("Serving %s on %s (tls)\n", path, addr)
		log.Printf("Using:\ncert: %s\nkey: %s\n", certFile, keyFile)
		if err := http.ListenAndServeTLS(addr, certFile, keyFile, nil); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Serving %s on %s\n", path, addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}
}
