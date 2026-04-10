package HelloWorld

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, There, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil) // Port is not 80 because, for example on Linux it need permisions for taking this port.
}
