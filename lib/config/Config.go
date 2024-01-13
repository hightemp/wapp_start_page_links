package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ConfigSite struct {
	Name        string
	Description string
	Image       string
	Url         string
}

type ConfigTheme struct {
	Name string
	Url  string
}

type ConfigSettings struct {
	Theme                string
	Themes               []ConfigTheme
	OpenLinksInNewWindow bool `yaml: "open_links_in_new_window"`
}

type ConfigStruct struct {
	List     []ConfigSite
	Settings ConfigSettings
}

type Config struct {
	Data     ConfigStruct
	FilePath string
}

func RemoveConfigSiteFromArray(slice []ConfigSite, s int) []ConfigSite {
	return append(slice[:s], slice[s+1:]...)
}

func (c *Config) CheckConfigFile() (bool, error) {
	if _, err := os.Stat(c.FilePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func (c *Config) InitDefaultConfig() {
	c.Data = ConfigStruct{
		Settings: ConfigSettings{
			Theme:                "light",
			OpenLinksInNewWindow: true,
		},
	}
}

func (c *Config) CreateDefaultConfig() {
	c.InitDefaultConfig()
	check, err := c.CheckConfigFile()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if !check {
		c.Save()
	}
}

func New(filePath string) *Config {
	config := &Config{
		FilePath: filePath,
	}
	config.CreateDefaultConfig()
	config.Load()
	return config
}

func (c *Config) Load() {
	data, err := os.ReadFile(c.FilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &c.Data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (c *Config) Save() {
	data, err := yaml.Marshal(c.Data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(c.FilePath, data, 0666)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
