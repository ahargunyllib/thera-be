package env

import (
	"time"

	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv             string        `mapstructure:"APP_ENV"`
	AppPort            string        `mapstructure:"APP_PORT"`
	APIKey             string        `mapstructure:"API_KEY"`
	DBHost             string        `mapstructure:"DB_HOST"`
	DBPort             string        `mapstructure:"DB_PORT"`
	DBUser             string        `mapstructure:"DB_USER"`
	DBPass             string        `mapstructure:"DB_PASS"`
	DBName             string        `mapstructure:"DB_NAME"`
	JwtSecretKey       string        `mapstructure:"JWT_SECRET_KEY"`
	JwtExpTime         time.Duration `mapstructure:"JWT_EXP_TIME"`
	RedisAddress       string        `mapstructure:"REDIS_ADDR"`
	RedisPassword      string        `mapstructure:"REDIS_PASSWORD"`
	GomailHost         string        `mapstructure:"GOMAIL_HOST"`
	GomailSenderEmail  string        `mapstructure:"GOMAIL_SENDER_EMAIL"`
	GomailSenderName   string        `mapstructure:"GOMAIL_SENDER_NAME"`
	GomailPort         int           `mapstructure:"GOMAIL_PORT"`
	GomailUsername     string        `mapstructure:"GOMAIL_USERNAME"`
	GomailPassword     string        `mapstructure:"GOMAIL_PASSWORD"`
	AWSAccessKeyID     string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	R2AccountID        string        `mapstructure:"R2_ACCOUNT_ID"`
	OpenAIAPIKey       string        `mapstructure:"OPENAI_API_KEY"`
}

var AppEnv = getEnv()

func getEnv() *Env {
	env := &Env{}

	viper.SetConfigFile("./config/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(log.CustomLogInfo{
			"error": err.Error(),
		}, "[ENV][getEnv] failed to read config file")
	}

	if err := viper.Unmarshal(env); err != nil {
		log.Fatal(log.CustomLogInfo{
			"error": err.Error(),
		}, "[ENV][getEnv] failed to unmarshal to struct")
	}

	switch env.AppEnv {
	case "development":
		log.Info(nil, "Application is running on development mode")
	case "production":
		log.Info(nil, "Application is running on production mode")
	case "staging":
		log.Info(nil, "Application is running on staging mode")
	}

	return env
}
