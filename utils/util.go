package utils

import (
	"log"
	"os"
	"strings"
)

// Get run time cwd
func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// Error handler
func ErrHadle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Clear file or directions
func CleanTmp(p string) {
	if p == "" {
		return
	}
	err := os.RemoveAll(p)
	if err != nil {
		log.Fatal(err)
	}
}

// Check command
func ValidCmd(cmd string) (string, string, bool) {
	cmds := strings.Split(cmd, " ")
	if len(cmds) > 1 && cmds[0] != "" && cmds[1] != "" {
		return cmds[0], strings.Join(cmds[1:], " "), true
	}
	return "", "", false
}
