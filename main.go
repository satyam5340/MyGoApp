package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Booking struct {
	Id      string `json:”id”`
	User    string `json:”user”`
	Members string `json:”members”`
}

var db *gorm.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Guys")
}
func createBooking(w http.ResponseWriter, r *http.Request) {
	//id := r.FormValue("id")
	user := r.FormValue("user")
	members := r.FormValue("members")

	booking := Booking{

		User:    user,
		Members: members,
	}

	// reqBody, _ := ioutil.ReadAll(r.Body)

	// fmt.Println(string(reqBody))
	// fmt.Println(json.Unmarshal([]byte(reqBody), &booking))
	// fmt.Println(booking)
	db.Create(&booking)
	fmt.Println("Endpoint Hit: Creating New Booking")
	json.NewEncoder(w).Encode(booking)
}

func handleRequest() {
	fmt.Println("Starting your developement server for testing api and orm")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/new-booking", createBooking).Methods("POST")
	http.ListenAndServe(":8000", router)

}
func main() {

	db, err = gorm.Open("mysql", "root:MyNewPassword@tcp(127.0.0.1:3306)/Football?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to open", err)
	} else {
		log.Println("Connection Established")
	}
	handleRequest()

	db.AutoMigrate(&Booking{})
}
