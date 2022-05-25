package Go_Server

import (
	"fmt"
	"net/http"
)

func FormFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Problem in Parsing %v", err)
		return
	}
	fmt.Fprintf(w, "Post request succesful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Address %s\n", address)

}
