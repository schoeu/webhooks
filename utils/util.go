package utils

import (
	"log"
	"os"
)

// get run time cwd
func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// error handler
func ErrHadle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// clear file or directions
func CleanTmp(p string) {
	if p == "" {
		return
	}
	err := os.RemoveAll(p)
	if err != nil {
		log.Fatal(err)
	}
}
