package main

import (
	"./config"
	"./utils"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
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
	cmd, para, vild := utils.ValidCmd(command)
	if vild {
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

		for k, v := range c.GetAll() {
			if k == action {
				execCmd(v)
				io.WriteString(w, v)
			}
		}
	})
}

// Exec command.
func execCmd(in string) {
	c, para, _ := utils.ValidCmd(in)
	cmd := exec.Command(c, para)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
