package configs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type configRaw struct {
	N         uint64  `yaml:"n"`
	P    string  `yaml:"p"`
}

type Config struct {
	configRaw
	P time.Duration
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cRaw configRaw

	err := unmarshal(&cRaw)
	if err != nil {
		return fmt.Errorf("YAML unmarshal: %w", err)
	}

	c.configRaw = cRaw


	c.P, err = time.ParseDuration(cRaw.P)
	if err != nil {
		return fmt.Errorf("closed_duration parse: %w", err)
	}

	return nil
}


func (c Config) Validate() error {
	if c.N != 0 {
		return errors.New("cannot be zero")
	}
	if c.P != 0 {
		return errors.New("cannot be zero")
	}

	return nil
}

func LoadConfig(configPath string) (Config, error) {
	configYAML, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("read config %s file: %w", configPath, err)
	}

	var c Config

	err = yaml.Unmarshal(configYAML, &c)
	if err != nil {
		return Config{}, fmt.Errorf("YAML unmarshal config: %w", err)
	}

	return c, nil
}
