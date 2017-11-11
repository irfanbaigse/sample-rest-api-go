package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/people", getPeople)
	fmt.Println("listening... 8192")
	err := http.ListenAndServe(":8192", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("INFO: No PORT environment variable detected, defaulting to 9192")

}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
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
