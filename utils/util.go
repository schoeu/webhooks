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
	c := strings.Split(cmd, " ")

	if len(c) > 1 && c[0] != "" && c[1] != "" {
		return c[0], strings.Join(c[1:], " "), true
	}
	return "", "", false
}

// analysis conf string.
func Analysis(s string) (string, string) {
	var router, cmd string
	cmds := strings.Split(s, ":")

	if len(cmds) < 2 || cmds[0] == "" || cmds[1] == "" {
		return router, cmd
	}
	return cmds[0], cmds[1]
}

func CmdFilter(c string) string {
	rs := strings.Replace(c, "sh ", "", -1)
	return rs
}
