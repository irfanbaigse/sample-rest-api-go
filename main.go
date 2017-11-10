package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	println("rest-api");
	router.HandleFunc("/people", getPeople).Methods("GET");
	fmt.Println("INFO: No PORT environment variable detected, defaulting to 9192")

	log.Fatal(http.ListenAndServe(":9192", router))
}

type Person struct {
	Id        int `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	people = append(people, Person{Id: 1, FirstName: "Irfan", Address: &Address{City: "Bangalore", State: "Karnantaka"}})
	people = append(people, Person{Id: 2, FirstName: "Imran", Address: &Address{City: "Mangalore", State: "Karnantaka"}})

	json.NewEncoder(w).Encode(people);
}
