package utils

import (
	"example.com/blogArch/gateway/configs"
	"golang.org/x/crypto/bcrypt"
	
	"log"
	"errors"
	"strings"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GrabEntries() []string {
	log.Println("In GrabEntries Util Function")
	query := "SELECT entry from entries WHERE \"user\" = 'user1';"
	rows, err := configs.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var res string
	allEntries := make([]string, 0)
	for rows.Next() {
		rows.Scan(&res)
		allEntries = append(allEntries, res)
		log.Printf("Entry is %s\n", res)
	}
	return allEntries
}

func InsertEntry(entry string, Username string) string {
	log.Println("In InsertEntry Util Function")
	entry = strings.Replace(entry, "\\'", "''", -1)
	query := "INSERT INTO entries (\"user\", entry) VALUES ('" + Username + "', '" + entry + "');"
	_, err := configs.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		return "Could not insert due to internal server error."
	}
	return "Inserted entry!"
}

func TryRegister(Username string, Password string) string {
	hash, _ := HashPassword(Password)
	// check database
	log.Println("In TryRegister Util Function")
	query := "INSERT INTO users (\"user\", password) SELECT '" + Username + "', '" + hash + "' WHERE NOT EXISTS (SELECT \"user\" FROM users WHERE \"user\" = '" + Username + "') LIMIT 1 RETURNING \"user\";"
	rows, err := configs.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var res string
	for rows.Next() {
		rows.Scan(&res)
	}
	return res
}

func TryLogin(Username string, Password string) (string, error) {
	log.Println("In TryLogin Util Function")
	query := "SELECT password from users WHERE \"user\" = '" + Username + "';"
	rows, err := configs.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer rows.Close()
	var res string
	for rows.Next() {
		rows.Scan(&res)
	}
	if res == "" || !CheckPasswordHash(Password, res) {
		return "Username or password does not match", errors.New("error in TryLogin")
	} else {
		return "Logged in", nil
	}
}