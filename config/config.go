package config

import (
	"bufio"
	"bytes"
	"github.com/schoeu/webhooks/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type configMap struct {
	value      map[string]string
	configPath string
}

// Set a key, value to map.
func (c *configMap) Set(k string, v string) {
	if c.value[k] == "" && k != "" {
		c.value[k] = v
	}
	// Store data to file when change the data.
	c.Store()
}

// Get value from configMap by key.
func (c *configMap) Get(k string) string {
	return c.value[k]
}

// Return config map.
func (c *configMap) GetAll() map[string]string {
	return c.value
}

// Store data to local file.
func (c *configMap) Store() {
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
func (c *configMap) Init(path string) {
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
func InitConfig(path string) configMap {
	c := configMap{}
	c.value = map[string]string{}
	c.Init(path)
	return c
}
