package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnHTTP(w http.ResponseWriter, code int, errText string) {
	w.WriteHeader(code)
	err := `{"msg": "` + errText + `"}`
	w.Write([]byte(err))
}

func getAllFood() []food {
	records, err := db.ReadAll("foods")
	if err != nil {
		log.Fatal("Couldnt get foods from db for initialisation")
	}
	var foods []food
	for _, f := range records {
		var food food
		if err := json.Unmarshal([]byte(f), &food); err != nil {
			log.Fatal("couldnt unmarshal food from db")
		}
		foods = append(foods, food)
	}
	return foods
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
