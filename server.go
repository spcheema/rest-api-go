package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "First route!")
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/hello", sayHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}