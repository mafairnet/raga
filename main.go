package main

import (
	"log"
	"net/http"
)

var configuration = getProgramConfiguration()

func main() {
	//command("database show AMPUSER")
	//Inicializo servidor web
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
