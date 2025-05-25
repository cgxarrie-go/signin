package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var lock = &sync.Mutex{}

var cfg *config

// Instance returns singleton instance.
func Instance() *config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()
		cfg = &config{
			Desks: make(Desks),
		}
	}

	return cfg
}

var (
	configFileName string = "signin.config"
)

type config struct {
	Bearer string
	Desks
	AttendanceFreeDays map[string]string
}

type Desks map[int]string

// Load read config file into config struct
func (c *config) Load() (err error) {
	fileName, err := c.fileName()
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); err != nil {
		return fmt.Errorf("configuration file not found\n" +
			"Use config command to add configuration")
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error loading config file: %s", err.Error())

	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return fmt.Errorf("Error parsing config: %s", err.Error())
	}

	return nil
}

// Save saves the config file
func (c *config) Save() (err error) {
	fileName, err := c.fileName()
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(b))

	return err
}

func (c *config) fileName() (folder string, err error) {
	ex, err := os.Executable()
	if err != nil {
		return folder, err
	}
	exPath := filepath.Dir(ex)
	fileName := fmt.Sprintf("%s/%s", exPath, configFileName)
	return fileName, nil
}
