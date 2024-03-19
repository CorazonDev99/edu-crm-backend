package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
	"log"
	"os"
	"sync"
	"time"
)

var (
	instance *Configuration
	once     sync.Once
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})
	return instance
}

// Configuration ...
type Configuration struct {
	AppName                    string
	AppVersion                 string
	AppURL                     string
	Environment                string
	ServerPort                 int
	ServerHost                 string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration
	// context timeout in seconds
	CtxTimeout                int
	SigninKey                 string
	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
	HashKey                   string
	PostgresHost              string
	PostgresPort              int
	PostgresDatabase          string
	PostgresUser              string
	PostgresPassword          string
	PostgresSSLMode           string
	MinioAccessKeyID          string
	MinioSecretKey            string
	MinioEndpoint             string
	MinioBucketName           string
	MinioLocation             string
	MinioUseSSL               bool
	DefaultOffset             string
	DefaultLimit              string
	DefaultPage               string
	DefaultPageSize           string
}

func load() *Configuration {
	return &Configuration{
		AppName:                   cast.ToString(getOrReturnDefault("APP_NAME", "EduCRM")),
		AppVersion:                cast.ToString(getOrReturnDefault("APP_VERSION", "1.0")),
		AppURL:                    cast.ToString(getOrReturnDefault("APP_URL", "localhost:7070")),
		ServerHost:                cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:                cast.ToInt(getOrReturnDefault("SERVER_PORT", "7070")),
		Environment:               cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		CtxTimeout:                cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		SigninKey:                 cast.ToString(getOrReturnDefault("SIGNING_KEY", "")),
		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
		PostgresHost:              cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:              cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase:          cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:              cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword:          cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		PostgresSSLMode:           cast.ToString(getOrReturnDefault("POSTGRES_SSLMODE", "disable")),
		MinioAccessKeyID:          cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "")),
		MinioSecretKey:            cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
		MinioEndpoint:             cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
		MinioBucketName:           cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
		MinioLocation:             cast.ToString(getOrReturnDefault("MINIO_LOCATION", "")),
		MinioUseSSL:               cast.ToBool(getOrReturnDefault("MINIO_USE_SSL", false)),
		DefaultOffset:             cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", 0)),
		DefaultLimit:              cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", 10)),
		HashKey:                   cast.ToString(getOrReturnDefault("HASH_KEY", "")),
		DefaultPage: cast.ToString(getOrReturnDefault(
			"DEFAULT_PAGE", 1)),
		DefaultPageSize: cast.ToString(getOrReturnDefault(
			"DEFAULT_PAGE_SIZE", 10)),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
