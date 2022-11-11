package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type greetingType struct {
	Message string
}
type people struct {
	User string
}

var greeting greetingType

func init() {
	greeting = greetingType{Message: "hello"}
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(greeting)
}
func fetchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := people{User: params["name"]}

	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", echoHello).Methods("GET")
	router.HandleFunc("/user/{name}", fetchUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
