package config

import (
	"../utils"
	"bufio"
	"bytes"
	"io"
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
	// Store data to file when change the data.
	c.Store()
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
			bf.WriteString(":")
			bf.WriteString(v)
			bf.WriteString("\n")
		}
	}

	// Clear exist confit file befor store it.
	// utils.CleanTmp(c.configPath)
	// Store data to local file.
	fi, err := os.OpenFile(c.configPath, os.O_WRONLY|os.O_APPEND, 0666)
	utils.ErrHadle(err)
	fi.Write([]byte(bf.String()))
	defer fi.Close()
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
	fi, err := os.Open(path)
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
		cmds := strings.Split(content, ":")
		if len(cmds) > 1 {
			c.value[cmds[0]] = cmds[1]
		}
	}
}

// Initial the config path.
func (c *ConfigMap) Init(path string) {
	if !filepath.IsAbs(path) {
		path = filepath.Join(utils.GetCwd(), path)
	}

	c.configPath = path

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}
			newFile.Close()
		} else {
			c.readConfig()
		}
	}
}

// Export config map.
func InitConfig(path string) ConfigMap {
	c := ConfigMap{}
	c.value = map[string]string{}
	c.Init(path)
	return c
}
