package config

import (
	"../utils"
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ConfigMap struct {
	value      map[string]string
	configPath string
}

var (
	DefaultConfPath = ".executor.conf"
)

// Set a key, value to map.
func (c *ConfigMap) Set(k string, v string) {
	if c.value[k] == "" && k != "" {
		c.value[k] = v
	}
}

// Get value from ConfigMap by key.
func (c *ConfigMap) Get(k string) string {
	return c.value[k]
}

// Return config map.
func (c *ConfigMap) GetAll() map[string]string {
	return c.value
}

// Store data to local file.
func (c *ConfigMap) Store() {
	var bf bytes.Buffer
	for k, v := range c.value {
		if k != "" {
			bf.WriteString(k)
			bf.WriteString(" ")
			bf.WriteString(v)
			bf.WriteString("\n")
		}
	}

	// Clear exist confit file befor store it.
	utils.CleanTmp(c.configPath)
	// Store data to local file.
	if e := ioutil.WriteFile(c.configPath, []byte(bf.String()), 0777); e != nil {
		log.Fatal(e)
	}
}

// Clear the map,
// just clear the object， not the config file content.
func (c *ConfigMap) Clear() {
	for k, _ := range c.value {
		delete(c.value, k)
	}
}

// Refresh config map data.
// Get newer data to ConfigMap
func (c *ConfigMap) Refresh() {
	// TODO: other process.
	c.readConfig()
}

// Read config file to struct
func (c *ConfigMap) readConfig() {
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

// Initial the config path.
func (c *ConfigMap) Init(path string) {
	if path == "" {
		cwd := utils.GetCwd()
		path = filepath.Join(cwd, DefaultConfPath)
	}

	c.configPath = path
}

// Export config map.
func InitConfig(path string) ConfigMap {
	c := ConfigMap{}
	c.value = map[string]string{}
	c.Init(path)
	return c
}
