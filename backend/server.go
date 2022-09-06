package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var counter = 0

	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		// Set cors config
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Max-Age", "15")

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, getStringFromNumber(counter))
		case "POST":
			counter += 1
			fmt.Fprintf(w, getStringFromNumber(counter))
			break
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getStringFromNumber(value int) string {
	// Declare variale and set number to string decimal
	i := strconv.FormatInt(int64(value), 10)

	return i
}
