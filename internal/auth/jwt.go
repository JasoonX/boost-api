package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/config"
)

const (
	defaultAccessExpiration  = 15 * time.Minute
	defaultRefreshExpiration = 24 * time.Hour
	subjectAccessToken       = "access"
	subjectRefreshToken      = "refresh"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Provider interface {
	GenerateTokenPair(info UserInfo) (*TokenPair, error)
}

type provider struct {
	JwtSecretKey []byte
	JwtIssuer    string
}

func New(cfg config.Config) Provider {
	return &provider{
		JwtSecretKey: cfg.SecretKey(),
		JwtIssuer:    cfg.Issuer(),
	}
}

func (p *provider) GenerateTokenPair(info UserInfo) (*TokenPair, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: info.Email,
		Role:  info.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   subjectAccessToken,
			ExpiresAt: time.Now().Add(defaultAccessExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    p.JwtIssuer,
		},
	})
	accessTokenSigned, err := accessToken.SignedString(p.JwtSecretKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate access token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: info.Email,
		Role:  info.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   subjectRefreshToken,
			ExpiresAt: time.Now().Add(defaultRefreshExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    p.JwtIssuer,
		},
	})
	refreshTokenSigned, err := refreshToken.SignedString(p.JwtSecretKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	return &TokenPair{
		AccessToken:  accessTokenSigned,
		RefreshToken: refreshTokenSigned,
	}, nil
}
