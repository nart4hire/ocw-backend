package env

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Environment struct {
	AppEnvironment string `env:"ENV"`
	ListenAddress  string `env:"LISTEN_ADDR" envDefault:"0.0.0.0"`

	ListenPort   int    `env:"PORT" envDefault:"8080"`
	LogtailToken string `env:"LOGTAIL_TOKEN"`

	HttpReqTimeout   int64 `env:"HTTP_TIMEOUT_SEC" envDefault:"1"`
	LogFlushInterval int64 `env:"LOG_FLUSH_INTERVAL_MS" envDefault:"1000"`

	UseReporter        bool   `env:"USE_REPORTER" envDefault:"true"`
	DatabaseConnection string `env:"DB_STRING"`

	PasswordCost int `env:"PASSWORD_COST" envDefault:"10"`

	TokenMethod         string `env:"TOKEN_SIGNING_METHOD" envDefault:"hs512"`
	TokenSecret         string `env:"TOKEN_SECRET"`
	TokenRefreshExpired int64  `env:"TOKEN_REFRESH_EXPIRED_MS" envDefault:"86400000"`
	TokenAccessExpired  int64  `env:"TOKEN_ACCESS_EXPIRED_MS" envDefault:"300000"`
	TokenIssuer         string `env:"TOKEN_ISSUER" envDefault:"ocw"`
}

func New() (*Environment, error) {
	if os.Getenv("ENV") == "PRODUCTION" {
		return NewEnv()
	}

	return NewDotEnv()
}

func NewEnv() (*Environment, error) {
	cfg := &Environment{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewDotEnv() (*Environment, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return NewEnv()
}
