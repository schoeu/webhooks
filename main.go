package main

import (
	"./config"
	"flag"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const (
	defaultConfPath = ".executor.conf"
	defaultPort     = ":8910"
)

func main() {
	var filePath, helper, port, command string
	flag.StringVar(&filePath, "path", defaultConfPath, "path of config file.")
	flag.StringVar(&port, "port", defaultPort, "server port.")
	flag.StringVar(&helper, "help", "", "help")
	flag.StringVar(&command, "add", "", "add router and command.")

	flag.Parse()

	// Get config instance.
	c := config.InitConfig(filePath)

	cmds := strings.Split(command, " ")
	if len(cmds) > 0 && cmds[0] != "" {
		c.Set(cmds[0], strings.Join(cmds[1:], " "))
	}

	router(c)
	http.ListenAndServe(defaultPort, nil)
}

func router(c config.ConfigMap) {
	urlReg := regexp.MustCompile("/tasks/(\\w+)$")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		info := urlReg.FindAllStringSubmatch(path, -1)
		action := ""
		if len(info) > 0 && len(info[0]) > 0 {
			action = info[0][1]
		}

		for k, v := range c.GetAll() {
			if k == action {
				io.WriteString(w, v)
			}
		}
	})
}

func exec(cmd string) {

}
