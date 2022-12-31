package utils

import (
	// "example.com/blogArch/gateway/configs"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TryRegister(Username string, Password string) string{
	// hash, _ := HashPassword(password)
	// check database
	fmt.Println("In TryRegister Util Function")
	connStr := "postgresql://newuser:password@localhost/blog?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// var q = "INSERT INTO users (\"user\", password) \n VALUES ('" + Username + "', '" + Password + "') \n RETURNING *;"
	q := "INSERT INTO users (\"user\", password) SELECT '" + Username + "', '" + Password + "' WHERE NOT EXISTS (SELECT \"user\" FROM users WHERE \"user\" = '" + Username + "') LIMIT 1 RETURNING \"user\";"
	log.Println(q)
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var res string
	for rows.Next() {
		rows.Scan(&res)
		fmt.Println(res)
	}
	return res
}

func TryLogin(Username string, Password string) string{
	// hash, _ := HashPassword(password)
	return ""
}