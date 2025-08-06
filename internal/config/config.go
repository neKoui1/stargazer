package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure: "database"`
	Server   ServerConfig   `mapstructure: "server"`
	App      AppConfig      `mapstructure: "app"`
}

type DatabaseConfig struct {
	Host     string `mapstructure: "host"`
	Port     string `mapstructure: "port"`
	Name     string `mapstructure: "name"`
	User     string `mapstructure: "user"`
	Password string `mapstructure: "password"`
	SSLMode  string `mapstructure: "ssl_mode"`
}

type ServerConfig struct {
	Port string `mapstructure: "port"`
	Host string `mapstructure: "host"`
}

type AppConfig struct {
	Name  string `mapstructure: "name"`
	Env   string `mapstructure: "env"`
	Debug bool   `mapstructure: "debug"`
}

func LoadConfig() (*Config, error) {

	envPath := "D:/stargazer/.env"
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		fmt.Printf("config: .env file does not exist at %s\n", envPath)
	}
	// 检查.env文件是否存在
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		fmt.Printf("Warning: .env file does not exist at %s\n", envPath)
	} else {
		fmt.Printf(".env file found at %s\n", envPath)
	}
	viper.SetConfigFile(envPath)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config: .env file not found: %v\n", err)
	}

	// fmt.Println("config: All keys from config: ")
	// for _, key := range viper.AllKeys() {
	// 	fmt.Printf(" %s = %v\n", key, viper.Get(key))
	// }
	// 手动设置配置值，确保正确的键名映射
	config := Config{
		Database: DatabaseConfig{
			Host:     viper.GetString("database_host"),
			Port:     viper.GetString("database_port"),
			Name:     viper.GetString("database_name"),
			User:     viper.GetString("database_user"),
			Password: viper.GetString("database_password"),
			SSLMode:  viper.GetString("database_ssl_mode"),
		},
		Server: ServerConfig{
			Host: viper.GetString("server_host"),
			Port: viper.GetString("server_port"),
		},
		App: AppConfig{
			Name:  viper.GetString("app_name"),
			Env:   viper.GetString("app_env"),
			Debug: viper.GetBool("app_debug"),
		},
	}
	setDefaults(&config)

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf(
			"config: failed to validate config: %w",
			err,
		)
	}

	return &config, nil
}

func setDefaults(config *Config) {
	if config.Database.Host == "" {
		config.Database.Host = "localhost"
	}
	if config.Database.Port == "" {
		config.Database.Port = "5432"
	}
	if config.Database.Name == "" {
		config.Database.Name = "stargazer"
	}
	if config.Database.SSLMode == "" {
		config.Database.SSLMode = "disable"
	}
	if config.Server.Port == "" {
		config.Server.Port = "8080"
	}
	if config.App.Name == "" {
		config.App.Name = "Stargazer"
	}
	if config.App.Env == "" {
		config.App.Env = "development"
	}
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.Database.Port == "" {
		return fmt.Errorf("database port is required")
	}
	if c.Database.Name == "" {
		return fmt.Errorf("database name is required")
	}
	if c.Database.User == "" {
		return fmt.Errorf("database user is required")
	}
	if c.Database.Password == "" {
		return fmt.Errorf("database password is required")
	}
	if c.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode)
}
