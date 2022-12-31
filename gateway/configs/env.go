package configs

import (
	"os"
)

func EnvUsername() string {
	token := os.Getenv("USER")
	return token
}

func EnvPassword() string {
	token := os.Getenv("PASSWORD")
	return token
}