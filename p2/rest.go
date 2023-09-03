// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type User struct {
// 	ID        int       `json:"id" bson:"_id"`
// 	Name      string    `json:"name" bson:"name"`
// 	Mobile    string    `json:"mobile" bson:"mobile"`
// 	Latitude  float64   `json:"latitude" bson:"latitude"`
// 	Longitude float64   `json:"longitude" bson:"longitude"`
// 	CreatedAt time.Time `json:"created_at" bson:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
// }

// var Users []User

// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var User User
// 	_ = json.NewDecoder(r.Body).Decode(&User)
// 	Users = append(Users, User)
// 	json.NewEncoder(w).Encode(User)
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/users", CreateMovie).Methods("POST")
// 	fmt.Printf("Starting server at post 8000\n")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// 	fmt.Println(Users)
// }
