package config

import (
	"runtime"

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

func Load(file string) (*config, error) {
	_, filename, _, _ := runtime.Caller(0)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filename + "/../../config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
