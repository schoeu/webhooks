package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schoeu/webhooks/exec"
)

type jsonStr struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Query   url.Values
}

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
		fmt.Println(err)
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

// Return json
func ReturnJson(w http.ResponseWriter, content string) {
	rs := jsonStr{}
	rs.Data = content
	rs.Status = 0
	js, err := json.Marshal(rs)
	if err != nil {
		rs.Status = 1
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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

// Exec command.
func ExecCmds(in string) string {
	output := exec.Run(strings.TrimSpace(in))
	a := string((*output).Stdout)
	e := (*output).Stderr
	if a == "" {
		a = "Task done."
	}
	if e != nil {
		a = fmt.Sprintf(" \"%v\"\n", e)
	}
	return a
}
