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

	MailingProvider string `env:"MAIL_PROVIDER" envDefault:"smtp"`
	MailingInterval int64  `env:"MAIL_INTERVAL_MS" envDefault:"1000"`

	SmtpIdentity string `env:"SMTP_IDENTITY"`
	SmtpUsername string `env:"SMTP_USERNAME"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
	SmtpServer   string `env:"SMTP_SERVER"`
	SmtpPort     int    `env:"SMTP_PORT" envDefault:"25"`
	SmtpAuthType string `env:"SMTP_TYPE" envDefault:"CRAM"`

	FrontendBaseURL   string `env:"FE_BASE_URL"`
	ResetPasswordPath string `env:"RESET_PASSWORD_PATH" envDefault:"/resetPassword"`

	EmailVerificationPath          string `env:"EMAIL_VERIFICATION_PATH" envDefault:"/verification"`
	EmailVerificationMaxRetry      int64  `env:"EMAIL_VERIFICATION_MAX_RETRY" envDefault:"5"`
	EmailVerificationRetryInterval int64  `env:"EMAIL_VERIFICATION_RESET_RETRY_INTERVAL_M" envDefault:"5"`
	EmailVerificationExpire        int64  `env:"EMAIL_VERIFICATION_EXPIRE_S" envDefault:"300"`

	RedisConnection string `env:"REDIS_STRING"`
	RedisPort       string `env:"REDIS_PORT" envDefault:"6379"`
	RedisUsername   string `env:"REDIS_USERNAME"`
	RedisPassword   string `env:"REDIS_PASSWORD"`
	RedisUseAuth    bool   `env:"REDIS_USE_AUTH" envDefault:"false"`
	RedisPrefixKey  string `env:"REDIS_PREFIX_KEY" envDefault:"app:"`
	RedisUseTLS     bool   `env:"REDIS_USE_TLS" envDefault:"false"`

	BucketEndpoint  string `env:"BUCKET_ENDPOINT"`
	BucketSecretKey string `env:"BUCKET_SECRET_KEY"`
	BucketAccessKey string `env:"BUCKET_ACCESS_KEY"`
	BucketTokenKey  string `env:"BUCKET_TOKEN_KEY"`
	BucketUseSSL    bool   `env:"BUCKET_USE_SSL" envDefault:"true"`
	BucketName      string `env:"BUCKET_NAME"`

	BucketSignedPutDuration int64 `env:"BUCKET_SIGNED_PUT_DURATION_S" envDefault:"36000"`
	BucketSignedGetDuration int64 `env:"BUCKET_SIGNED_GET_DURATION_S" envDefault:"1800"`

	BucketMaterialBasePath string `env:"BUCKET_MATERIAL_BASE_PATH" envDefault:"materials"`

	UseBucket bool `env:"USE_BUCKET" envDefault:"true"`
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
	err := godotenv.Load(
		".env",
		".env.local",
	)

	if err != nil {
		return nil, err
	}

	return NewEnv()
}
