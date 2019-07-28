package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DBServer    string
	DBPort      string
	Database    string
	DbUser      string
	DbPassword  string
	LogDir      string
	LogLevel    string
	ServicePort string
	UrlService  string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if len(os.Args) < 2 {
		panic("Please input the config file dir for the service")
	}
	comfigFile := os.Args[1]
	fmt.Println("Service config file in: ", comfigFile)
	if _, err := toml.DecodeFile(comfigFile, &c); err != nil {
		fmt.Println(err)
	}
}
