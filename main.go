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
	r.HandleFunc("/food/{name}", returnByName)
	r.HandleFunc("/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	os.Setenv("ACCESS_SECRET", "asdffgsdfgh")
	db = initDB()
	genUserMap()
	log.Println("arfid encyclopedia server started on port 8080")
	log.Print(getAllUsers())
	//log.Printf("admin user login: %v %v", users[1].Username, users[1].PasswordHash)
	foods = []food{
		food{Name: "Vegimite", Category: "Condiments", Visual: "description", Texture: "description", Smell: "description", Taste: "description", Nutrients: []string{"smth"}},
		food{Name: "Mozzarella", Category: "Dairy", Visual: "White, soft but solid (it will deform a bit when you poke it, but you can’t spread it)", Texture: "feels wet, slightly chewy, slightly stringy", Smell: "Very weak/non existent smell", Taste: "Weak, slightly milky flavour, some brands can have an unpleasant bitter/acidic aftertaste, though this might go away when cooked", Nutrients: []string{"fats"}},
	}
	handleRequests()
}
