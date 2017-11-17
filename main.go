package main

import (
	"./config"
	"./exec"
	"./utils"
	"flag"
	"io"
	"net/http"
	"regexp"
)

const (
	defaultConfPath = ".webhook.conf"
	defaultPort     = ":8910"
)

func main() {
	var filePath, helper, port, command string
	flag.StringVar(&filePath, "path", defaultConfPath, "path of config file.")
	flag.StringVar(&port, "port", defaultPort, "server port.")
	flag.StringVar(&helper, "help", "", "help")
	flag.StringVar(&command, "add", "", "add router and command.")

	flag.Parse()

	// Usage like this:  webhook --add "some_router:echo test_string"
	// then request `localhost:8910/some_router` to run the command `echo test_string`

	// Get config instance.
	c := config.InitConfig(filePath)
	cmd, para := utils.Analysis(command)
	if cmd != "" && para != "" {
		c.Set(cmd, para)
	}

	router(c)
	http.ListenAndServe(defaultPort, nil)
}

// Router for actions.
func router(c config.ConfigMap) {
	urlReg := regexp.MustCompile("/tasks/(\\w+)$")
	cmdReg := regexp.MustCompile("/run/(.+)$")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		runInfo := cmdReg.FindAllStringSubmatch(path, -1)
		var action, runCmd string
		hit := false

		if len(runInfo) > 0 && len(runInfo[0]) > 0 {
			runCmd = runInfo[0][1]
			//o := execCmd(runCmd, "run")
			rs := execCmds(runCmd)
			hit = true
			io.WriteString(w, rs)
			return
		}

		info := urlReg.FindAllStringSubmatch(path, -1)
		if len(info) > 0 && len(info[0]) > 0 {
			action = info[0][1]
		}
		allActions := c.GetAll()
		for k, v := range allActions {
			if k == action {
				rs := execCmds(v)
				hit = true
				io.WriteString(w, rs)
			}
		}
		if !hit {
			io.WriteString(w, "No action for this router")
		}
	})
}

// Exec command.
func execCmds(in string) string {
	output := exec.Series(in)
	a := string((*output[0]).Stdout)
	e := (*output[0]).Stderr
	utils.ErrHadle(e)
	if a == "" {
		a = "Task done."
	}
	return a
}
