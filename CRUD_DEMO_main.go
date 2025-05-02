package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Capitals struct {
	Country string `json:"country"`
	Caps    string `json:"caps"`
}

var caps []Capitals

func getCaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caps)

}

func CreateCaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cap Capitals
	err := json.NewDecoder(r.Body).Decode(&cap)
	if err != nil {
		http.Error(w, "Decode not ok", http.StatusNoContent)
		return
	}

	caps = append(caps, cap)

	json.NewEncoder(w).Encode(caps)

}

func GetCap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range caps {

		if item.Country == params["country"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	http.Error(w, "Not Found", http.StatusNotFound)

}

func DelCap(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range caps {
		if item.Country == params["country"] {
			caps = append(caps[:index], caps[index+1:]...)
			break
		}
		http.Error(w, "Not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(caps)

}

func PutCap(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, item := range caps {
		if item.Country == params["country"] {
			caps = append(caps[:index], caps[index+1:]...)
			var cap Capitals
			error := json.NewDecoder(r.Body).Decode(&cap)
			if error != nil {
				return
			}
			cap.Country = params["country"]
			caps = append(caps, cap)
			break
		}
	}
	json.NewEncoder(w).Encode(caps)
}

func main() {

	caps = append(caps, Capitals{Country: "India", Caps: "New Delhi"})
	// object for new gorilla router
	r := mux.NewRouter()
	r.HandleFunc("/Details", getCaps).Methods("GET")
	r.HandleFunc("/Details", CreateCaps).Methods("POST")
	r.HandleFunc("/Details/{country}", GetCap).Methods("GET")
	r.HandleFunc("/Details/{country}", DelCap).Methods("DELETE")
	r.HandleFunc("/Details/{country}", PutCap).Methods("PUT")

	http.ListenAndServe(":8080", r)
}
