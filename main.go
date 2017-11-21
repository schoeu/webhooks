package main

import (
	"flag"
	"io"
	"net/http"
	"regexp"
	"os"

	"github.com/schoeu/webhooks/config"
	"github.com/schoeu/webhooks/routers"
	"github.com/schoeu/webhooks/utils"
)

const (
	defaultConfPath = ".webhook.conf"
	defaultPort     = "8910"
)

func main() {
	var filePath, helper, port, command, token string
	flag.StringVar(&filePath, "path", defaultConfPath, "path of config file.")
	flag.StringVar(&port, "port", defaultPort, "server port.")
	flag.StringVar(&helper, "help", "", "help")
	flag.StringVar(&command, "add", "", "add router and command.")
	flag.StringVar(&token, "token", "", "token for request")

	flag.Parse()

	// Usage like this:  webhook --add "some_router:echo test_string"
	// then request `localhost:8910/some_router` to run the command `echo test_string`

	// Get config instance.
	c := config.InitConfig(filePath)

	// Add restart command to refresh configuration
	args := os.Args
	if len(args) > 1 && args[1] == "restart" {
		c.Refresh()
		return
	}

	router(c, token)

	if command == "" {
		// Get command from configuration.
		c.Refresh()
		http.ListenAndServe(":"+port, nil)
	} else {
		cmd, para := utils.Analysis(command)
		if cmd != "" && para != "" {
			c.Set(cmd, para)
		}
	}
}

// Router for actions.
func router(c config.ConfigMap, token string) {
	urlReg := regexp.MustCompile("/tasks/(\\w+)$")
	cmdReg := regexp.MustCompile("/run/(.+)$")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/" {
			io.WriteString(w, "Server is ready.")
			return
		}

		// Token check.
		if token != "" {
			valid := utils.CheckToken(r, token)
			if !valid {
				io.WriteString(w, "Wrong token.")
				return
			}
		}

		runInfo := cmdReg.FindAllStringSubmatch(path, -1)
		taskInfo := urlReg.FindAllStringSubmatch(path, -1)

		hit := false

		if len(runInfo) > 0 && len(runInfo[0]) > 0 {
			routers.RunRouter(w, runInfo)
			return
		}

		if len(taskInfo) > 0 && len(taskInfo[0]) > 0 {
			hit = routers.TaskRouter(w, taskInfo, c)
		}

		if !hit {
			io.WriteString(w, "No action for this router")
		}
	})
}
