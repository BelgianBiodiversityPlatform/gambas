package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GBIF!")
}

func main() {
	listenAddressPtr := flag.String("listen", "127.0.0.1:8080", "The address the server will listen to.")

	flag.Parse()

	http.HandleFunc("/v1", handler)

	fmt.Printf("Gambas is now listening on %s\n", *listenAddressPtr)
	http.ListenAndServe(*listenAddressPtr, nil)
}
