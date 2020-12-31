package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Students struct {
	Name string
	Age  int
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World!")

}

func getStudents(w http.ResponseWriter, r *http.Request) {

	newStudent := []Students{
		{Name: "rambabu", Age: 35},
		{Name: "avyan", Age: 7},
		{Name: "Sree Lakshmi", Age: 30},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(newStudent)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newStudent)

	}

}
func main() {

	//multiplexer
	http.HandleFunc("/greet", greet)

	http.HandleFunc("/students", getStudents)

	//server listen
	http.ListenAndServe("localhost:8008", nil)

}
