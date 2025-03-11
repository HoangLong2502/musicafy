package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	MigrationURL        string `mapstructure:"MIGRATION_URL"`
	TOKEN_SYMMETRIC_KEY string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ZINGMP3_URL         string `mapstructure:"ZINGMP3_URL"`
	ZINGMP3_AC_URL      string `mapstructure:"ZINGMP3_AC_URL"`
	ZINGMP3_VERSION     string `mapstructure:"ZINGMP3_VERSION"`
	ZINGMP3_API_KEY     string `mapstructure:"ZINGMP3_API_KEY"`
	MAIL_SENDER         string `mapstructure:"EMAIL_SENDER_NAME"`
	MAIL_PASSWORD       string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	MAIL_ADDRESS        string `mapstructure:"EMAIL_SENDER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		err = nil
		return
	}

	err = viper.Unmarshal(&config)
	return
}
