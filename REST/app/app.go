package app

import (
	"net/http"
)

func Start() {
	//multiplexer

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", Greet)

	mux.HandleFunc("/students", GetStudents)

	//server listen
	http.ListenAndServe("localhost:8008", mux)

}
