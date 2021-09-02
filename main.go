package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/common"
	"go-mongo-rest-api/handlers"
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
	router.HandleFunc("/member", handlers.MemberHandler).Methods("POST", "PUT")
	router.HandleFunc("/member/{email}", handlers.MemberHandler).Methods("GET", "DELETE")
	router.HandleFunc("/member/find/", handlers.FindMemberHandler).Methods("GET")
	router.HandleFunc("/member/find/{skills}&{title}", handlers.FindMemberHandler).Methods("GET")
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
