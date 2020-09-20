package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfg *configuration
)

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	cfg = &configuration{}
	return decoder.Decode(cfg)
}

func DB() db {
	return *cfg.DB
}

func Mode() string {
	return *cfg.Mode
}

func Log() log {
	return *cfg.Log
}

func Server() server {
	return *cfg.Server
}

func Email() email {
	return *cfg.Email
}

type configuration struct {
	DB   *db     `yaml:"db"`
	Mode *string `yaml:"mode"`
	Log  *log    `yaml:"log"`

	Server *server `yaml:"server"`
	Email  *email  `yaml:"email"`
}

type db struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
}

type log struct {
	Path string `yaml:"path"`
}

type server struct {
	Port       string `yaml:"port"`
	AssetsPath string `yaml:"assets_path"`
	Secret     string `yaml:"secret"`
}

type email struct {
	Token string `json:"token"`
}
