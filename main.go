package main

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

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

	query := "INSERT INTO users (email, password) values ($1, $2)"
	for {
		email, password := email_password()

		// Hashed Password Generation
		password = hashAndSlat([]byte(password))
		_, err := db.Query(query, email, password)
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
			email, password := email_password()
			var passwordForValidate string
			queryForValidate := "SELECT  password FROM users WHERE email=$1"
			err := db.QueryRow(queryForValidate, email).Scan(&passwordForValidate)
			if err != nil {
				log.Info("Error", err)
				return
			}
			//Hashed Password Compare
			matched, err := CompareHashPassword(passwordForValidate, []byte(password))
			if err != nil {
				return
			}
			if matched {
				fmt.Println("Successfully Logged in ", email)
			} else {
				fmt.Println("Not Matched")
			}
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

func email_password() (string, string) {
	var email, password string
	fmt.Println("Please Enter Email")
	_, err := fmt.Scan(&email)
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
	return email, password
}

func hashAndSlat(password []byte) string {
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Info("Error : ", err)
		return ""
	}
	return string(hashPassword)
}

func CompareHashPassword(hashedPassword string, password []byte) (bool, error) {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		log.Info("Error : ", err)
		return false, err
	}
	return true, nil
}
