package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Read the contents of the HTML file
	htmlFile, err := ioutil.ReadFile("page.html")
	if err != nil {
		http.Error(w, "Unable to load HTML file", http.StatusInternalServerError)
		return
	}

	// Set the content type header to "text/html"
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML response
	w.Write(htmlFile)
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

