package service

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Println("Starting HTTP service at " + port)

	router := NewRouter()

	http.Handle("/", router)

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
