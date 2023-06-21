package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Host      string `yaml:"host" mapstructure:"host"`
	Port      int    `yaml:"port" mapstructure:"port"`
	Username  string `yaml:"username" mapstructure:"username"`
	Password  string `yaml:"password" mapstructure:"password"`
	DbName    string `yaml:"db_name" mapstructure:"db_name"`
	Timezone  string `yaml:"timezone" mapstructure:"timezone"`
	EnableSsl bool   `yaml:"enable_ssl" mapstructure:"enable_ssl"`
}

type config struct {
	Database        dbConfig `yaml:"database" mapstructure:"database"`
	Port            int      `yaml:"port" mapstructure:"port"`
	JwtSecret       string   `yaml:"jwt_secret" mapstructure:"jwt_secret"`
	SwaggerBasePath string   `yaml:"swagger_base_path" mapstructure:"swagger_base_path"`
}

func Load(file string) (*config, error) {
	_, filename, _, _ := runtime.Caller(0)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filename + "/../../config")

	log.Println("Loading config file: ", file+".yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	log.Println("Config: ", cfg)
	log.Println("DB Config: ", cfg.Database)
	return cfg, nil
}
