package utils

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB

func DBSetup() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "test"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	// TODO: セットでどこかで db.Close() はしなくていいのか調査が必要
	gormDB, err := gorm.Open(DBMS, CONNECT)
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

	db = gormDB
}

func DB() *gorm.DB {
	return db
}

func GetAllUsers(db *gorm.DB) *gorm.DB {
	var users []User
	result := db.Order("created_at asc").Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result
}

func GetUser(db *gorm.DB, userId string) *gorm.DB {
	var user User
	result := db.First(&user, userId)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result
}

func CreateUser(db *gorm.DB, user *User) *gorm.DB {
	data := User{
		Name:  user.Name,
		Email: user.Email,
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	result := db.Create(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result
}

func UpdateUser(db *gorm.DB, userId string, user *User) *gorm.DB {
	data := User{
		Name:  user.Name,
		Email: user.Email,
	}
	data.UpdatedAt = time.Now()
	result := db.Model(&User{}).Where("id = ?", userId).Updates(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result
}

func DeleteUser(db *gorm.DB, userId string) *gorm.DB {
	var user User
	result := db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result
}
