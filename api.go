package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type food struct {
	Name      string   `json:"Name"`
	Category  string   `json:"Category"`
	Visual    string   `json:"Visual"`
	Texture   string   `json:"Texture"`
	Smell     string   `json:"Smell"`
	Taste     string   `json:"Taste"`
	Nutrients []string `json:"Nutrients"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the ARFID encyclopedia api root")
	log.Println("endpoint hit: root")
}

func returnAll(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint hit: returnAll")
	json.NewEncoder(w).Encode(foods)
}

func returnByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("endpoint hit: food " + name)
	for _, food := range foods {
		if food.Name == name {
			json.NewEncoder(w).Encode(food)
		}
	}
}

func createNewFood(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	log.Printf("token string: %v", token)
	userid, err := verifyToken(token)
	if err != nil {
		log.Fatal(err)
	}
	if users[userid].AccessLevel >= MOD {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var food food
		json.Unmarshal(reqBody, &food)
		log.Println("endpoint hit: create food " + food.Name)
		foods = append(foods, food)
		json.NewEncoder(w).Encode(food)
	} else {
		fmt.Fprintf(w, `{"Error": "Not authorised for this endpoint"}`)
	}

}
