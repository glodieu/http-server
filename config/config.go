package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"time"
)

type Config struct {
	Listen       string        `yaml:"listen"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

func ParseConfigFile(file string) (*Config, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Config file cannot be found: %v\n", err))
	}
	defer fp.Close()
	return ParseConfig(fp)
}

func ParseConfig(r io.Reader) (*Config, error) {
	config := &Config{}

	decoder := yaml.NewDecoder(r)
	decoder.SetStrict(true)
	err := decoder.Decode(&config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid yaml config: %v\n", err))
	}

	return config, nil
}
