package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}
type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags:= flag.String("config","","path to the configuration file")
		flag.Parse()

		configPath= *flags
	}

	if configPath ==""{
		log.Fatal("Config path is not set")
	}

	if _,error :=os.Stat(configPath); os.IsNotExist(error) {
		log.Fatalf("config didn't exist: %s",configPath)

	}

	var cfg Config
	error:=cleanenv.ReadConfig(configPath,&cfg)
	if error !=nil{
		log.Fatalf("can not read config file: %s",error.Error())
	}

	return &cfg
}