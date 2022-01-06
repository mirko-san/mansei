package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mirko-san/mansei/app/src/middleware"
	"github.com/mirko-san/mansei/app/src/utils"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	utils.DBSetup()
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	db := utils.DB()
	result := utils.GetAllUsers(db)
	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func returnUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	db := utils.DB()
	result := utils.GetUser(db, userId)
	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user utils.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db := utils.DB()
	result := utils.CreateUser(db, &user)
	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	body, _ := ioutil.ReadAll(r.Body)
	var user utils.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db := utils.DB()
	result := utils.UpdateUser(db, userId, &user)
	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	db := utils.DB()
	result := utils.DeleteUser(db, userId)
	responseBody, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.Handle("/users", middleware.EnsureValidToken()(http.HandlerFunc(createUser))).Methods("POST")
	myRouter.Handle("/users", middleware.EnsureValidToken()(http.HandlerFunc(returnAllUsers)))
	myRouter.Handle("/users/{userId}", middleware.EnsureValidToken()(http.HandlerFunc(updateUser))).Methods("PUT")
	myRouter.Handle("/users/{userId}", middleware.EnsureValidToken()(http.HandlerFunc(deleteUser))).Methods("DELETE")
	myRouter.Handle("/users/{userId}", middleware.EnsureValidToken()(http.HandlerFunc(returnUser)))
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
