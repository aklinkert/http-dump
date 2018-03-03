package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var (
	port = flag.Int("port", 8080, "Port")
	verbose = flag.Bool("verbose", false, "Print all reqests to stdout")
)

func main() {
	flag.Parse()

	log.Println("Starting server on port", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.HandlerFunc(handle)); err != nil {
		log.Fatal(err)
	}
}

func handle (w http.ResponseWriter, r *http.Request) {
		content, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}

		if *verbose == true {
			log.Printf("Request from %v: %s", r.RemoteAddr, content)
		}

		w.Write(content)
		fmt.Fprintf(w, "RemoteAddr: %v\n", r.RemoteAddr)
}
