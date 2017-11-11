package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/people", getPeople)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
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
