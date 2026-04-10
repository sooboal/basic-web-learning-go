package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
		Dynamic aspect: http.Request contains all info about request and it's parameters
		You can read GET params with `r.URL.Query().Get("token")` or POST params ith r.FormValue("email")
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!, %s\n", r.URL.Query().Get("token")) // If u want to check token u need to go by this link: http://localhost:8080/?token=abc123
	})

	fs := http.FileServer(http.Dir("static/")) // Says: "Files(on disk) are stored in ./static directory according to program execution"
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	/*
		Why do we need http.StripPrefix here? Because without it this will happen:
		Query: http://localhost:8080/static/style.css
		Server will try to access at: static/static/style.css
	*/

	http.ListenAndServe(":8080", nil)
	/*
		To check result of this step go to http://localhost:8080/static/index.html.
		But Make sure you in the right Directory
	*/
}
