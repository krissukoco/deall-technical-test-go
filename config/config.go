package config

import (
	"os"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	DbName    string `yaml:"db_name"`
	Timezone  string `yaml:"timezone"`
	EnableSsl bool   `yaml:"enable_ssl"`
}

type config struct {
	Database  *dbConfig `yaml:"database"`
	Port      int       `yaml:"port"`
	JwtSecret string    `yaml:"jwt_secret"`
}

var Config *config

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Config = &config{}
	err = viper.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
}
