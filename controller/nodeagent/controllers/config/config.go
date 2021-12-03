package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func FromFile(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if nil != err {
		panic(err)
	}
	cf := &Config{}
	err = yaml.Unmarshal(file, cf)
	if nil != err {
		panic(fmt.Errorf("Failed to parse config from config file! Error:%s", err.Error()))
	}

	return cf
}
