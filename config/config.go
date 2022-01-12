package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	Instance *Config
)

type Config struct {
	path     string
	AppKey   string `json:"app_key"`
	SecKey   string `json:"sec_key"`
	From     string `json:"from"`
	To       string `json:"to"`
	SignType string `json:"sign_type"`
}

func Init(path string) {
	c := &Config{
		path: path,
	}
	if err := c.load(); err != nil {
		color.Red(err.Error())
		color.Green("Create a new configuration in %v", path)
	}

	c.save()
	Instance = c
}

// load from path
func (c *Config) load() (err error) {
	file, err := os.Open(c.path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, c)
}

// save file to path
func (c *Config) save() (err error) {
	var data bytes.Buffer
	encoder := json.NewEncoder(&data)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(c)
	if err == nil {
		os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
		err = ioutil.WriteFile(c.path, data.Bytes(), 0644)
	}
	if err != nil {
		color.Red("Cannot save config to %v\n%v", c.path, err.Error())
	}
	return
}
