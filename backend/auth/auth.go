package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	j      = os.Getenv("JWT_SECRET")
	secret = []byte(strings.TrimSpace(j))
)

// GenerateToken generates a JWT token and returns the token and any errors.
//
//   - minutesToExpire: The number of minutes before the token expires.
//
// The token is signed using the secret key.
//
// Error can be nil.
func GenerateToken(minutesToExpire int, customClaims jwt.MapClaims) (string, error) {
	claims := jwt.MapClaims{
		"exp": jwt.NewNumericDate(time.Now().Add(time.Duration(minutesToExpire) * time.Minute)),
		"iss": "github.com/joevtap",
		"iat": jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	for key, value := range customClaims {
		token.Claims.(jwt.MapClaims)[key] = value
	}

	return token.SignedString(secret)
}

// ParseToken parses a JWT token and returns the token and any errors.
//
// The token is verified using the secret key.
//
// Error can be nil.
func ParseToken(signedString string) (*jwt.Token, error) {
	return jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
}

// GetTokenClaims returns the claims of a JWT token.
func GetTokenClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}

// IsTokenExpiring returns true if the token is expiring in the next minutesLeft minutes.
func IsTokenExpiring(token *jwt.Token, minutesLeft int) bool {
	claims := GetTokenClaims(token)
	expirationTime := claims["exp"]

	return time.Until(time.Unix(int64(expirationTime.(float64)), 0)).Minutes() <= float64(minutesLeft)
}
