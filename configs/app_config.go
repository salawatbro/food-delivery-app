package configs

import "github.com/spf13/viper"

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
	JwtConfig      JwtConfig
}

type AppConfig struct {
	Name    string
	Version string
	Port    string
	Host    string
	Env     string
}

func LoadConfig() Config {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("config")
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	config := Config{
		AppConfig: AppConfig{
			Name:    v.GetString("app.name"),
			Version: v.GetString("app.version"),
			Port:    v.GetString("app.port"),
			Host:    v.GetString("app.host"),
			Env:     v.GetString("app.env"),
		},
		DatabaseConfig: DatabaseConfig{
			Host:     v.GetString("database.host"),
			Port:     v.GetString("database.port"),
			Username: v.GetString("database.user"),
			Password: v.GetString("database.password"),
			DBName:   v.GetString("database.name"),
		},
		JwtConfig: JwtConfig{
			SecretKey: v.GetString("jwt.secret"),
			ExpiresIn: v.GetDuration("jwt.expires_in"),
		},
	}

	return config
}
