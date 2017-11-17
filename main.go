package main

import (
	"./config"
	"./utils"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os/exec"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		info := urlReg.FindAllStringSubmatch(path, -1)
		action := ""
		if len(info) > 0 && len(info[0]) > 0 {
			action = info[0][1]
		}
		hit := false
		allActions := c.GetAll()
		for k, v := range allActions {
			if k == action {
				o := execCmd(v)
				hit = true
				io.WriteString(w, string(o))

			}
		}
		if !hit {
			io.WriteString(w, "No action for this path: "+action)
		}
	})
}

// Exec command.
func execCmd(in string) []byte {
	in = utils.CmdFilter(in)
	c, para, ok := utils.ValidCmd(in)
	if !ok {
		return []byte("")
	}
	cmd := exec.Command(c, para)

	o, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	return o
}
