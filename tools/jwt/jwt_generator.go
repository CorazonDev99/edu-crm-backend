package jwt

import (
	"strconv"
	"strings"
	"time"

	"EduCRM/config"

	jwt "github.com/dgrijalva/jwt-go"
)

var conf = config.Config()

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

// GenerateNewTokens handler_func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id, role, roleTitle string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id, role, roleTitle)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken(id)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id, roleID, roleTitle string) (string, error) {

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expire"] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()
	claims["roleID"] = roleID
	claims["roleTitle"] = roleTitle

	// Set private token credentials:
	// for _, credential := range credentials {
	// 	claims[credential] = true
	// }

	// in local server access token ttl = 10 days
	if conf.Environment == "develop" {
		claims["expire"] = time.Now().Add(time.Minute * time.Duration(100*conf.JWTSecretKeyExpireMinutes)).Unix()
	} else {
		// in staging server access token ttl = day
		claims["expire"] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()
	}
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(conf.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken(id string) (string, error) {
	// Create a new SHA256 hash.

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expire"] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()

	// Set private token credentials:
	// for _, credential := range credentials {
	// 	claims[credential] = true
	// }

	// in local server access token ttl = 10 days
	if conf.Environment == "develop" {
		claims["expire"] = time.Now().Add(time.Minute * time.Duration(100*conf.JWTSecretKeyExpireMinutes)).Unix()
	} else {
		// in staging server access token ttl = day
		claims["expire"] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()
	}
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(conf.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

// ParseRefreshToken handler_func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
