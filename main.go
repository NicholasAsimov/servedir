package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		host string
		port int
		dir  string
		ssl  bool
		cert string
		key  string
	)

	flag.StringVar(&host, "host", "localhost", "host to listen on")
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.StringVar(&dir, "dir", ".", "directory to serve")
	flag.BoolVar(&ssl, "ssl", false, "use ssl")
	flag.StringVar(&cert, "cert", "", "path to ssl certificate, requires -ssl")
	flag.StringVar(&key, "key", "", "path to ssl key, requires -ssl")

	flag.Parse()

	if ssl && (cert == "" || key == "") {
		log.Println("Error: Specify ssl cert and key paths")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !ssl && (cert != "" || key != "") {
		log.Println("Error: Using -cert or -key requires -ssl")
		flag.PrintDefaults()
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))
	
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening on %s...", addr)

	if ssl {
		log.Fatal(http.ListenAndServeTLS(addr, cert, key, nil))
	} else {
		log.Fatal(http.ListenAndServe(addr, nil))
	}
}
