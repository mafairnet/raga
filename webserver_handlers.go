package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RAGA | REST Asterisk Golang Agent, please look the API documentation for the supported requests!\n")
}

func Command(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	commandString := vars["command"]
	data := command(commandString)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
