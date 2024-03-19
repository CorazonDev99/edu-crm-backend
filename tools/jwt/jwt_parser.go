package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID        string
	UserRole      string
	UserRoleTitle string
	Expires       int64
}

// ExtractTokenMetadata handler_func to extract metadata from JWT.
func ExtractTokenMetadata(ctx *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, userIDExist := claims["id"]
		if !userIDExist {
			return nil, errors.New("unauthorized")
		}
		userRole, roleIdExist := claims["roleID"]
		if !roleIdExist {
			return nil, errors.New("unauthorized")
		}
		userRoleTitle, roleTitleExist := claims["roleTitle"]
		if !roleTitleExist {
			return nil, errors.New("unauthorized")
		}
		// Expires time.
		expires, expiresExist := claims["expire"]
		if !expiresExist {
			return nil, errors.New("unauthorized")
		}
		expiresInt64 := int64(expires.(float64))
		if expiresInt64 < time.Now().Unix() {
			return nil, errors.New("token expired")
		}
		// User credentials.
		return &TokenMetadata{
			UserID:        userID.(string),
			Expires:       expiresInt64,
			UserRole:      userRole.(string),
			UserRoleTitle: userRoleTitle.(string),
		}, nil
	}
	return nil, err
}
func ExtractRefreshTokenMetadata(ctx *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, userIDExist := claims["id"]
		if !userIDExist {
			return nil, errors.New("unauthorized")
		}

		// Expires time.
		expires, expiresExist := claims["expire"]
		if !expiresExist {
			return nil, errors.New("unauthorized")
		}
		expiresInt64 := int64(expires.(float64))
		if expiresInt64 < time.Now().Unix() {
			return nil, errors.New("token expired")
		}
		// User credentials.
		return &TokenMetadata{
			UserID:  userID.(string),
			Expires: expiresInt64,
		}, nil
	}
	return nil, err
}

func extractToken(ctx *gin.Context) string {
	bearToken := ctx.GetHeader("Authorization")
	token := fmt.Sprintf("%v", bearToken)
	onlyToken := strings.Split(token, " ")
	if len(onlyToken) != 2 {
		return ""
	}
	return onlyToken[1]
}

func verifyToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(conf.JWTSecretKey), nil
}
