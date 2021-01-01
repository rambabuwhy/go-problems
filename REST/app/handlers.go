package app

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

func Greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World!")

}

func GetStudents(w http.ResponseWriter, r *http.Request) {

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
