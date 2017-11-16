package main

import (
	"net/http"
	"flag"
	"./config"
)

const (
	defaultConfPath = ".executor.conf"
	defaultPort = ":8913"
)

func main() {
	var filePath, helper, port string
	flag.StringVar(&filePath, "path", defaultConfPath, "path of config file.")
	flag.StringVar(&port, "port", defaultPort, "server port.")
	flag.StringVar(&helper, "help", "", "help")

	flag.Parse()

	if filePath == "" {
		initConf()
	}

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request){

	})
	http.ListenAndServe(defaultPort, nil)
}

func router() {

}