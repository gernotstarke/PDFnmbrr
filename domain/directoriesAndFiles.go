package domain

import (
	"log"
	"os"
)

func GetUserHomeDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func GetOSUserName() string {
	return "anonymous user"
}
