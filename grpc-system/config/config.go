package config

import (
	srv_db "github.com/heyue7/BookHub-Server-Learn/service/srv-db"
	srv_grpc "github.com/heyue7/BookHub-Server-Learn/service/srv-grpc"
	"gopkg.in/yaml.v2"
	"os"
)

var Conf Config

type Config struct {
	Env string `yaml:"env"`

	Server srv_grpc.Config `yaml:"server"`

	DB srv_db.DBConfig `yaml:"db"`
}

func Init(yamlFile string) (err error) {
	if yamlFile == "" {
		yamlFile = ".yaml"
	}

	buf, err := os.ReadFile(yamlFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, &Conf)

	return
}
