package configuration

import (
	"log"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/yaml"
)

const configFile = "config.yml"
const envVarsPrefix = "MYAPP"

type datebase struct {
	Host     string `yaml:"host" cfgRequired:"true"`
	Port     int    `yaml:"port" cfgRequired:"true"`
	Dbname   string `yaml:"dbname" cfgRequired:"true"`
	User     string `yaml:"user" cfgRequired:"true"`
	Password string `yaml:"password" cfgRequired:"true"`
}

type bcrypt struct {
	Cost int `yaml:"cost" cfgDefault:"14"`
}

type jwt struct {
	SecretKey string `yaml:"secretkey" cfgDefault:"changeit"`
	ExpireIn  int64  `yaml:"expires_in" cfgDefault:"600"` // Expiration in second
}

type configApp struct {
	Database datebase `yaml:"database"`
	Bcrypt   bcrypt   `yaml:"bcrypt"`
	JWT      jwt      `yaml:"jwt"`
}

var App configApp

func LoadConfig() {

	App = configApp{}

	goconfig.FileEnv = "CONFIG_FILE_NAME"
	goconfig.PathEnv = "CONFIG_FILE_PATH"

	goconfig.File = configFile
	goconfig.PrefixEnv = envVarsPrefix
	err := goconfig.Parse(&App)
	if err != nil {
		log.Fatal(err)
	}

}
