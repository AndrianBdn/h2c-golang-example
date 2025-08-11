package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Configure protocols for H2C-only support (HTTP/2 cleartext)
	protocols := new(http.Protocols)
	protocols.SetUnencryptedHTTP2(true) // Enable H2C (HTTP/2 cleartext)
	protocols.SetHTTP1(false)           // Explicitly disable HTTP/1.1
	protocols.SetHTTP2(false)           // Explicitly disable encrypted HTTP/2 (HTTPS)

	// Note: You can modify the protocol support as needed:
	// - protocols.SetHTTP1(true) to allow HTTP/1.1 connections
	// - protocols.SetHTTP2(true) to allow encrypted HTTP/2 over TLS

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s (HTTP/%d.%d)", r.Method, r.URL.Path, r.ProtoMajor, r.ProtoMinor)
		fmt.Fprintf(w, "Hello from %s!\n", r.URL.Path)
		fmt.Fprintf(w, "Protocol: HTTP/%d.%d\n", r.ProtoMajor, r.ProtoMinor)
		fmt.Fprintf(w, "TLS: %v\n", r.TLS != nil)
	})

	server := &http.Server{
		Addr:      ":8080",
		Handler:   handler,
		Protocols: protocols,
	}

	log.Println("Starting H2C server on :8080")
	log.Println("Test with: curl -v --http2-prior-knowledge http://localhost:8080/test")
	log.Fatal(server.ListenAndServe())
}