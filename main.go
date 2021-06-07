package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/SomeshSunariwal/Password_Security_Hash_Method/config"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf(`user=%s host=%s dbname=%s password=%s sslmode=%s`,
		config.DBUser, config.DBHost, config.DBName, config.DBPassword, config.DBsslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	query := "INSERT INTO users (user_name, password) values ($1, $2)"
	for {
		userName, password := userName_password()
		_, err := db.Query(query, userName, password)
		if err != nil {
			log.Info("Error", err)
			return
		}
		var wantToValidate string
		fmt.Println("Want to Validate. Please write Yes/no")
		_, err = fmt.Scan(&wantToValidate)
		if err != nil {
			log.Info("Error : ", err)
			return
		}
		if strings.ToLower(wantToValidate) == "yes" {
			user_name, passwordForValidate := userName_password()
			queryForValidate := "SELECT user_name, password FROM users WHERE user_name=$1 and password=$2"
			err := db.QueryRow(queryForValidate, user_name, passwordForValidate).Scan(&user_name, &passwordForValidate)
			if err != nil {
				log.Info("Error", err)
				return
			}
			fmt.Println(user_name, passwordForValidate)
		}
		var WantToExit string
		fmt.Println("Want to Exit. Please write Yes/no")
		_, err = fmt.Scan(&WantToExit)
		if err != nil {
			log.Info("Error : ", err)
			return
		}
		if strings.ToLower(WantToExit) == "yes" {
			break
		}
	}
}

func userName_password() (string, string) {
	var userName, password string
	fmt.Println("Please Enter User Name")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Info("Error : ", err)
		return "", ""
	}
	fmt.Println("Please Enter Password")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Info("Error : ", err)
		return "", ""
	}
	return userName, password
}
