package main

import (
  "fmt"
	"encoding/json"
  "time"
	"log"
	"net/http"

	"github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  gorm.Model
  Name string
  Email string
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
  db := sqlConnect()
	db.AutoMigrate(&User{})
  defer db.Close()

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	db := sqlConnect()
	var users []User
	db.Order("created_at asc").Find(&users)
	defer db.Close()
	responseBody, err := json.Marshal(users)
	if err != nil {
			log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	db := sqlConnect()
	name := "name"
	email :=  "email"
	db.Create(&User{Name: name, Email: email})
	defer db.Close()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func sqlConnect() (database *gorm.DB) {
  DBMS := "mysql"
  USER := "root"
  PASS := "root"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "test"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

  count := 0
  db, err := gorm.Open(DBMS, CONNECT)
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      fmt.Print(".")
      time.Sleep(time.Second)
      count++
      if count > 10 {
        fmt.Println("")
        fmt.Println("DB接続失敗")
        panic(err)
      }
      db, err = gorm.Open(DBMS, CONNECT)
    }
  }
  fmt.Println("DB接続成功")

  return db
}
