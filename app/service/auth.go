package service

import (
	"demo/app/domain"
	"demo/app/repository"
	"demo/app/serializers"
	"demo/app/utils/methods"
	"demo/infra/config"
	"demo/infra/errors"
	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	Login(req *serializers.LoginRequest) (*serializers.LoginResponse, error)
	Logout(user *serializers.LoggedInUser) error
	//RefreshToken(refreshToken string) (*serializers.LoginResp, error)
	//VerifyToken(accessToken string) (*serializers.VerifyTokenResp, error)
}

type auth struct {
	userRepo     repository.IUsers
	tokenService IToken
}

func NewAuthService(userRepo repository.IUsers, tokenService IToken) IAuth {
	return &auth{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (as *auth) Login(req *serializers.LoginRequest) (*serializers.LoginResponse, error) {
	var user *domain.User
	var err error

	if user, err = as.userRepo.FindByEmail(req.Email); err != nil {
		return nil, errors.ErrInvalidEmail
	}

	loginPass := []byte(req.Password)
	hashedPass := []byte(*user.Password)

	if err = bcrypt.CompareHashAndPassword(hashedPass, loginPass); err != nil {
		return nil, errors.ErrInvalidPassword
	}

	var token *serializers.JwtToken

	if token, err = as.tokenService.CreateToken(user.ID); err != nil {
		return nil, errors.ErrCreateJwt
	}

	if err = as.tokenService.StoreTokenUuid(user.ID, token); err != nil {
		return nil, errors.ErrStoreTokenUuid
	}

	var userResp *serializers.UserResponse
	respErr := methods.StructToStruct(user, &userResp)

	if respErr != nil {
		return nil, respErr
	}

	res := &serializers.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User:         userResp,
	}
	return res, nil
}

func (as *auth) Logout(user *serializers.LoggedInUser) error {
	return as.tokenService.DeleteTokenUuid(
		config.Redis().AccessUuidPrefix+user.AccessUuid,
		config.Redis().RefreshUuidPrefix+user.RefreshUuid,
	)
}
