package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func LsdIndex(w http.ResponseWriter, r *http.Request) {
	projects := Projects{
		Project{Name: "Corc"},
		Project{Name: "AOSBR"},
	}

	if err := json.NewEncoder(w).Encode(projects); err != nil {
		panic(err)
	}
}

func LsdShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lsdId := vars["lsdId"]
	fmt.Fprintln(w, "LSD show:", lsdId)
}

