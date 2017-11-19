package utils

import (
	"../exec"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Get run time cwd
func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
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

// Token checker
func CheckToken(r *http.Request, token string) bool {
	u, err := url.ParseQuery(r.URL.RawQuery)
	ErrHadle(err)
	if u.Get("token") == token {
		return true
	}
	return false
}

// Exec command.
func ExecCmds(in string) string {
	output := exec.Series(strings.TrimSpace(in))
	a := string((*output[0]).Stdout)
	e := (*output[0]).Stderr
	if a == "" {
		a = "Task done."
	}
	if e != nil {
		a = fmt.Sprintf(" \"%v\"\n", e)
	}
	return a
}
