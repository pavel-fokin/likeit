package apiutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenDuration = time.Hour
	accessTokenIssuer   = "likeit-backend"
	accessTokenSubject  = "access"
	accessTokenAudience = "likeit-frontend"
)

// JWTClaims represents the claims of a JSON Web Token (JWT).
type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userID"`
}

// NewAccessToken generates a new access token for the given user ID.
// The access token is a JSON Web Token (JWT) signed with the HS256 algorithm.
// It includes the issuer, subject, audience, expiration time, and the user ID as claims.
// The access token is signed using a signing key.
// It returns the generated access token as a string and any error encountered during the process.
func NewAccessToken(userID string, signingKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    accessTokenIssuer,
			Subject:   accessTokenSubject,
			Audience:  jwt.ClaimStrings{accessTokenAudience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyAccessToken verifies the given access token.
// It returns the claims of the access token and any error encountered during the process.
func VerifyAccessToken(tokenString string, signingKey string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the signing method.
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// Verify the claims.
	if claims.Issuer != accessTokenIssuer {
		return nil, fmt.Errorf("unexpected issuer: %v", claims.Issuer)
	}
	if claims.Subject != accessTokenSubject {
		return nil, fmt.Errorf("unexpected subject: %v", claims.Subject)
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	// TODO: Verify the audience.

	return claims, nil
}
