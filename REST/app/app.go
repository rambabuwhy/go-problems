package app

func start() {
	//multiplexer

	mux := http.NewServerMux()

	http.HandleFunc("/greet", greet)

	http.HandleFunc("/students", getStudents)

	//server listen
	http.ListenAndServe("localhost:8008", nil)

}
