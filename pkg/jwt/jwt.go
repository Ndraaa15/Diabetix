package jwt

import (
	"net/http"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

var (
	ErrInvalidTokenExpired = errx.New().WithCode(iris.StatusUnauthorized).WithMessage("invalid token expired")
	ErrFailedClaimJWT      = errx.New().WithCode(iris.StatusUnauthorized).WithMessage("failed claim jwt")
	ErrInvalidSignature    = errx.New().WithCode(http.StatusUnauthorized).WithMessage("invalid signature")
)

func EncodeToken(user *domain.User, duration time.Duration) (string, error) {
	claims := &JWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "diabetix",
			Subject:   "authentication",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		return "", ErrInvalidTokenExpired.WithError(err)
	}
	return signedToken, nil
}

func DecodeToken(token string) (*JWTClaims, error) {
	decoded, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		token.Method = jwt.SigningMethodHS256
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if err != nil {
		return &JWTClaims{}, ErrInvalidSignature
	}

	if !decoded.Valid {
		return &JWTClaims{}, ErrInvalidTokenExpired
	}

	claims, ok := decoded.Claims.(*JWTClaims)
	if !ok {
		return &JWTClaims{}, ErrFailedClaimJWT
	}

	return claims, nil
}
