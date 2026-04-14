// Firstly we need to install package `mux` for routing
// go mod init 03-Routing
// go get -u github.com/gorilla/mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating album")
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting album")
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting album")
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating album")
}

func albumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Album handler")
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secure handler")
}

func insecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insecure handler")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	// URL Parameters
	// For example we have URL: /album/812/song/BaoBab
	r.HandleFunc("/album/{title}/song/{songName}", func(w http.ResponseWriter, r *http.Request) {
		// Get Album and go the song
		vars := mux.Vars(r)
		title := vars["title"]
		song := vars["songName"]
		fmt.Fprintf(w, "Album: %s, Song: %s\n", title, song)
	})

	// Also we can restrict the request handler to specific HTTP methods
	r.HandleFunc("/album/{title}", CreateAlbum).Methods("POST")
	r.HandleFunc("/album/{title}", GetAlbum).Methods("GET")
	r.HandleFunc("/album/{title}", UpdateAlbum).Methods("PUT")
	r.HandleFunc("/album/{title}", DeleteAlbum).Methods("DELETE")

	// Also we can restrict the request handler to specific hostnames
	r.HandleFunc("/album/{title}", albumHandler).Host("api.example.com")

	// Restrict the request handler to http/https
	r.HandleFunc("/secure", secureHandler).Schemes("https")
	r.HandleFunc("/insecure", insecureHandler).Schemes("http")

	// Restrict the request handler to specific prefixes

	albumRouter := r.PathPrefix("/album").Subrouter()
	albumRouter.HandleFunc("/{title}", GetAlbum).Methods("GET")
	albumRouter.HandleFunc("/{title}", UpdateAlbum).Methods("PUT")
	albumRouter.HandleFunc("/{title}", DeleteAlbum).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
