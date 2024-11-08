package env

import "github.com/spf13/viper"

type Env struct {
	DatabaseHost        string `mapstructure:"DATABASE_HOST"`
	DatabasePort        int    `mapstructure:"DATABASE_PORT"`
	DatabaseName        string `mapstructure:"DATABASE_NAME"`
	DatabaseUser        string `mapstructure:"DATABASE_USER"`
	DatabasePassword    string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseSSLMode     string `mapstructure:"DATABASE_SSL_MODE"`
	SnowFlakeNode       int64  `mapstructure:"SNOWFLAKE_NODE"`
	EmailHost           string `mapstructure:"EMAIL_HOST"`
	EmailPort           int    `mapstructure:"EMAIL_PORT"`
	EmailSender         string `mapstructure:"EMAIL_SENDER"`
	EmailPassword       string `mapstructure:"EMAIL_PASSWORD"`
	HtmlPath            string `mapstructure:"HTML_PATH"`
	GeminiApiKey        string `mapstructure:"GEMINI_API_KEY"`
	GeminiModel         string `mapstructure:"GEMINI_MODEL"`
	CloudinaryName      string `mapstructure:"CLOUDINARY_NAME"`
	CloudinaryApiKey    string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryFolder    string `mapstructure:"CLOUDINARY_FOLDER"`
	AppName             string `mapstructure:"APP_NAME"`
	AppAddr             string `mapstructure:"APP_ADDR"`
	AppPort             string `mapstructure:"APP_PORT"`
}

func New() *Env {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(".env")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var env Env
	v.Unmarshal(&env)

	return &env
}
