package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
// convert it to an integer using the strconv.Atoi() function. If it can't
// be converted to an integer, or the value is less than 1, we return a 404 page // not found response.
	id , err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "Display  specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		// use the http.Error method to send a non 200 status code and a plain error text response body  
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Set("Content-Type", "application/json")

		// w.Write([]byte(`{"name":"Alex"}`))
		//set a new cache-control header . If an existing "Cache-control" header exists 
		w.Header().Set("Cache-Control","public, max-age=31536000")

		// adds a new "Cache-Control header"
		w.Header().Add("Cache-Control","public")
		w.Header().Add("Cache-Control","public, max-age=31536000")

		// delete all values for the Cache-Control header
		w.Header().Del("Cache-Control")
		
		//Retrieve th first value for the Cache-Control heade
		w.Header().Get("Cache-Control")

		//Retrieves a slice of all values for the Cache-Control header
		w.Header().Values("Cache-Control")
		// w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
		
		//delete system generated headers 
		w.Header()["Date"] = nil 
		http.Error(w, "method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a specific snippet"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
