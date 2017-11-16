package config

import (
	"../utils"
	"path/filepath"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"io"
	"strings"
)

type configMap struct {
	value map[string]string
	configPath string
}

var (
	DefaultConfPath = ".executor.conf"
)

// Set a key, value to map.
func (c *configMap) Set(k string, v string) {
	if c.value[k] == "" {
		c.value[k] = v
	}
}

// Get value from configMap by key.
func (c *configMap) Get(k string) string {
	return c.value[k]
}

// Store data to local file.
func (c *configMap) Store() {
	var bf bytes.Buffer
	for k, v := range c.value {
		if k != "" {
			bf.WriteString(k)
			bf.WriteString(" ")
			bf.WriteString(v)
			bf.WriteString("\n")
		}
	}

	cwd := utils.GetCwd()
	path := filepath.Join(cwd, DefaultConfPath)
	c.configPath = path
	// Clear exist confit file befor store it.
	utils.CleanTmp(path)
	// Store data to local file.
	if e := ioutil.WriteFile(path, []byte(bf), 0777); e != nil {
		log.Fatal(e)
	}
}

// Clear the map,
// just clear the object， not the config file content.
func (c *configMap) Clear() {
	for k, _ := range c.value {
		delete(c.value, k)
	}
}

// Refresh config map data.
// Get newer data to configMap
func (c *configMap) Refresh() {
	// TODO: other process.
	c.readConfig()
}

// Read config file to struct
func (c *configMap) readConfig() {
	path := c.configPath
	fi, err := os.Open(filepath.Join(path))
	utils.ErrHadle(err)
	defer fi.Close()
	br := bufio.NewReader(fi)

	// write config content before clear it.
	c.Clear()

	for {
		a, _, co := br.ReadLine()
		if co == io.EOF {
			break
		}
		content := string(a)
		cmds := strings.Split(content, " ")
		if len(cmds) > 2 {
			c.value[cmds[0]] = strings.Join(cmds[1:], " ")
		}
	}
}
