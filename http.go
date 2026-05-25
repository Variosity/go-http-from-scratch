package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// The Hello Handler
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}
	// The JSON Handler
	jsonHandler := func(j http.ResponseWriter, jeq *http.Request) {
		// JSON Struct
		type jsonGroup struct {
			ID      int
			Name    string
			Message []string
		}
		// JSON Contents
		jsonr := jsonGroup{
			ID:      1,
			Name:    "Hello",
			Message: []string{"JSON", "JSON"},
		}
		// Set the content type
		j.Header().Set("Content-Type", "application/json")
		j.WriteHeader(http.StatusOK)
		// Parse the JSON
		b, err := json.Marshal(jsonr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			http.Error(j, err.Error(), http.StatusInternalServerError)
			return
		}
		// Convert the byte slice into a string
		s := string(b[:])
		io.WriteString(j, s)
	}
	// Registering the handler functions
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
