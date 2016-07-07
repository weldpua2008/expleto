package config

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	// "reflect"
	// "strconv"
	// "strings"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// Config stores configurations values
type Config struct {
	AppName   string `json:"app_name" yaml:"app_name" toml:"app_name"`
	BaseURL   string `json:"base_url" yaml:"base_url" toml:"base_url"`
	Port      int    `json:"port" yaml:"port" toml:"port"`
	Verbose   bool   `json:"verbose" yaml:"verbose" toml:"verbose"`
	StaticDir string `json:"static_dir" yaml:"static_dir" toml:"static_dir"`
	ViewsDir  string `json:"view_dir" yaml:"view_dir" toml:"view_dir"`
}

// DefaultConfig returns the default configuration settings.
func DefaultConfig() *Config {
	return &Config{
		AppName:   "expleto web app",
		BaseURL:   "http://localhost:9000",
		Port:      9000,
		Verbose:   false,
		StaticDir: "static",
		ViewsDir:  "views",
	}
}

// NewConfig reads configuration from path. The format is deduced from the file extension
//	* .json    - is decoded as json
//	* .yml     - is decoded as yaml
//	* .toml    - is decoded as toml
func NewConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	switch filepath.Ext(path) {
	case ".json":
		jerr := json.Unmarshal(data, cfg)
		if jerr != nil {
			return nil, jerr
		}
	case ".toml":
		_, terr := toml.Decode(string(data), cfg)
		if terr != nil {
			return nil, terr
		}
	case ".yml":
		yerr := yaml.Unmarshal(data, cfg)
		if yerr != nil {
			return nil, yerr
		}

	default:
		return nil, errors.New("utron: config file format not supported")
	}
	return cfg, nil
}