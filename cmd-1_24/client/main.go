package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get URL from command line or use default
	url := "http://localhost:8080/test"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	// Configure protocols for H2C-only support (HTTP/2 cleartext)
	protocols := new(http.Protocols)
	protocols.SetUnencryptedHTTP2(true) // Enable H2C (HTTP/2 cleartext)
	protocols.SetHTTP1(false)           // Explicitly disable HTTP/1.1
	protocols.SetHTTP2(false)           // Explicitly disable encrypted HTTP/2 (HTTPS)

	// Create client with H2C transport configuration
	client := &http.Client{
		Transport: &http.Transport{
			Protocols: protocols,
		},
	}

	// Make the request
	log.Printf("Requesting: %s", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Display response information
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Protocol: HTTP/%d.%d\n", resp.ProtoMajor, resp.ProtoMinor)
	fmt.Printf("Headers:\n")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	fmt.Printf("\nBody:\n%s", body)
}