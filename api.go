package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type food struct {
	Name      string   `json:"Name"`
	Author    string   `json: "Author"`
	Category  string   `json:"Category"`
	Visual    string   `json:"Visual"`
	Texture   string   `json:"Texture"`
	Smell     string   `json:"Smell"`
	Taste     string   `json:"Taste"`
	Nutrients []string `json:"Main Nutrient(s)"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	returnHTTP(w, 200, "Home endpoint")
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
	userid, err := verifyToken(token)
	checkErr(err)
	var user user
	if err := db.Read("users", userid, &user); err != nil {
		log.Println("Error user not found", err)
		returnHTTP(w, 500, "User not found")
		return
	}
	if user.AccessLevel >= MOD {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var food food
		json.Unmarshal(reqBody, &food)
		log.Println("endpoint hit: create food " + food.Name)
		foods = append(foods, food)
		if err := writeFoodToDB(food); err != nil {
			returnHTTP(w, 500, "writing food to db")
			log.Println("Error writing food to db: ", err)
			return
		}
		json.NewEncoder(w).Encode(food)
	} else {
		returnHTTP(w, 500, "User not authorised to access this endpoity")
	}

}

func editFood(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorisation"), "Bearer ")
	userid, err := verifyToken(token)
	checkErr(err)
	var user user
	if err := db.Read("users", userid, &user); err != nil {
		log.Println("Error User not found")
		returnHTTP(w, 500, "User not found")
	}
	if user.AccessLevel >= MOD {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var reqFood food
		json.Unmarshal(reqBody, &reqFood)
		var dbFood food
		if err := db.Read("foods", reqFood.Name, &dbFood); err != nil {
			log.Printf("Food %s not found", reqFood.Name)
			returnHTTP(w, 500, "Food "+reqFood.Name+" not found")
			return
		}
		finalFood := food{
			Name:      dbFood.Name,
			Category:  reqFood.Category,
			Visual:    reqFood.Visual,
			Texture:   reqFood.Texture,
			Smell:     reqFood.Smell,
			Taste:     reqFood.Taste,
			Nutrients: reqFood.Nutrients,
			Author:    reqFood.Author,
		}
		db.Write("foods", finalFood.Name, finalFood)
		returnHTTP(w, 200, "Edit successfull")
	} else {
		returnHTTP(w, 500, "User is not authorised to access this endpoints")
	}
}
