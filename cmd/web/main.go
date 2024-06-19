package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/* BE AWARE OF THIS
 * setting the content type in the header to match the fact that its JSON
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
*/
