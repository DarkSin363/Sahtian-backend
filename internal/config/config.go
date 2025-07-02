package config

import (
	"github.com/BigDwarf/sahtian/internal/log"

	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectId    string         `mapstructure:"project_id"`
	EnableDocs   bool           `mapstructure:"enable_docs"`
	Logger       *log.Config    `mapstructure:"logger"`
	Server       ServerConfig   `mapstructure:"http_server"`
	Database     DatabaseConfig `mapstructure:"database"`
	Telegram     TelegramConfig `mapstructure:"telegram"`
	Storage      StorageConfig  `mapstructure:"storage"`
	Debug        DebugConfig    `mapstructure:"debug"`
	PubSubConfig PubSubConfig   `mapstructure:"pubsub"`
	Features     Features       `mapstructure:"features"`
}

type Features struct {
	PaidAnswers bool `mapstructure:"paid_answers"`
}

type PubSubConfig struct {
	Topic string `json:"topic"`
}

type DebugConfig struct {
	Enabled bool  `mapstructure:"enabled"`
	UserId  int64 `mapstructure:"user_id"`
}

type StorageConfig struct {
	Bucket      string `mapstructure:"bucket"`
	Credentials string `mapstructure:"credentials"`
}

type CorsConfig struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
	AllowMethods []string `mapstructure:"allow_methods"`
}

type ServerConfig struct {
	Listen     string     `mapstructure:"listen"`
	CorsConfig CorsConfig `mapstructure:"cors"`
}

type DatabaseConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
}

type TelegramConfig struct {
	Token  string `mapstructure:"token"`
	AppUrl string `mapstructure:"app_url"`
}

func Init(cfgName string) (*Config, error) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.SetConfigName("default.yml")
		viper.AddConfigPath("./config")
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("SERVICE")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.WithStack(err)
	}

	conf := &Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.WithStack(err)
	}

	return conf, nil
}
