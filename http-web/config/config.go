package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Conf Config

type Config struct {
	Env string `yaml:"env"`

	Server struct {
		Name string `yaml:"name"`
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	Services struct {
		GrpcSystem string `yaml:"grpc-system"`
	} `yaml:"services"`
}

func Init(yamlFile string) error {
	if yamlFile == "" {
		yamlFile = ".yaml"
	}

	buf, err := os.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &Conf)
	return err
}
