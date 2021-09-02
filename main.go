package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/common"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func startServer() {
	common.Init()
	handleRequests()
}

func main() {
	fmt.Println("Server starting ...")
	startServer()
}
