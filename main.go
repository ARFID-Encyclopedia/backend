package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sdomino/scribble"

	"github.com/gorilla/mux"
)

var foods []food
var userMap map[string]string
var db *scribble.Driver

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/foods", returnAll)
	r.HandleFunc("/food", createNewFood).Methods("POST")
	r.HandleFunc("/food/edit", editFood).Methods("POST")
	r.HandleFunc("/food/{name}", returnByName)
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
	log.Fatal(http.ListenAndServe(":6969", r))
}

func main() {
	os.Setenv("ACCESS_SECRET", "asdffgsdfgh")
	db = initDB()
	genUserMap()
	log.Println("arfid encyclopedia server started on port 8080")
	foods = getAllFood()
	handleRequests()
}
