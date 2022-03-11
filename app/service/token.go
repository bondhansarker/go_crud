package service

import (
	"time"

	"demo/app/repository"
	"demo/app/serializers"
	"demo/infra/config"
	"demo/infra/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type IToken interface {
	CreateToken(userID uint) (*serializers.JwtToken, error)
	//StoreTokenUuid(userID uint, token *serializers.JwtToken) error
	//DeleteTokenUuid(uuid ...string) error
}

type token struct {
	userRepo repository.IUsers
}

func NewTokenService(userRepo repository.IUsers) IToken {
	return &token{
		userRepo: userRepo,
	}
}

func (t *token) CreateToken(userID uint) (*serializers.JwtToken, error) {
	var err error
	jwtConf := config.Jwt()
	token := &serializers.JwtToken{}

	token.AccessExpiry = time.Now().Add(time.Minute * jwtConf.AccessTokenExpiry).Unix()
	token.AccessUuid = uuid.New().String()

	token.RefreshExpiry = time.Now().Add(time.Minute * jwtConf.RefreshTokenExpiry).Unix()
	token.RefreshUuid = uuid.New().String()

	user, getErr := t.userRepo.Find(int(userID))
	if getErr != nil {
		return nil, errors.NewError(getErr.Message)
	}

	atClaims := jwt.MapClaims{}
	atClaims["uid"] = user.ID
	atClaims["aid"] = token.AccessUuid
	atClaims["rid"] = token.RefreshUuid
	atClaims["exp"] = token.AccessExpiry

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(jwtConf.AccessTokenSecret))
	if err != nil {
		return nil, errors.ErrAccessTokenSign
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["uid"] = user.ID
	rtClaims["aid"] = token.AccessUuid
	rtClaims["rid"] = token.RefreshUuid
	rtClaims["exp"] = token.RefreshExpiry

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(jwtConf.RefreshTokenSecret))
	if err != nil {
		return nil, errors.ErrRefreshTokenSign
	}
	return token, nil
}
