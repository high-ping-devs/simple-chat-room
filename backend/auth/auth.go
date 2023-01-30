package auth

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

// The same as GenerateToken but with a default expiration time of 15 minutes.
//
// Args can be used to override the default expiration time.
//
//   - args[0]: The number of minutes before the token expires.
func GenerateRefreshToken(customClaims jwt.MapClaims, args ...int) (string, error) {
	minutesToExpire := 15

	if len(args) > 0 {
		minutesToExpire = args[0]
	}

	return GenerateToken(minutesToExpire, customClaims)
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

// GrantPermission grants a permission to a token.
func GrantPermission(token *jwt.Token, permission string) {
	if !HasPermission(token, permission) {
		claims := GetTokenClaims(token)
		claims["permissions"] = append(claims["permissions"].([]string), permission)
	}
}

// RevokePermission revokes a permission from a token.
func RevokePermission(token *jwt.Token, permission string) {
	if HasPermission(token, permission) {
		claims := GetTokenClaims(token)
		permissions := claims["permissions"].([]string)

		for i, p := range permissions {
			if p == permission {
				permissions = append(permissions[:i], permissions[i+1:]...)
				break
			}
		}

		claims["permissions"] = permissions
	}
}

// HasPermission returns true if the token has the permission.
func HasPermission(token *jwt.Token, permission string) bool {
	claims := GetTokenClaims(token)
	permissions := claims["permissions"].([]string)

	for _, p := range permissions {
		if p == permission {
			return true
		}
	}

	return false
}

// HashAndSaltPassword hashes and salts a password.
//
// It returns the hash in string format.
func HashAndSaltPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}

	return string(hash)
}

// PasswordMatch returns true if the password matches the hash.
func PasswordMatch(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
