package authRoutes

import (
	"bybu/go-postgres/shared/config"
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"context"
	"fmt"
	"time"
)

type _IAuthService struct {
	userRepo user.IUserRepository
}

func NewUserAuthService() *_IAuthService{
	return &_IAuthService{
		userRepo: user.NewUserRepository(),
	}
}

func (service *_IAuthService) signup(request user.IRegisterRequest, ctx *context.Context) (user.IAuthUserResponse, models.IError) {
	hashPassword, err := utils.HashPassword(request.Password); if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "password is invalid" };
	}
	request.Password = hashPassword;
	newUser, userErr := service.userRepo.Create(request.CreateUser(), ctx);
	if (userErr != models.IError{}) {
		return user.IAuthUserResponse{}, userErr
	}

	accessToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: newUser.ID.String(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (tokenErr != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	
	refreshToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: newUser.ID.String(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (tokenErr != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create refresh token" }
	}

	return user.IAuthUserResponse{
		ID: newUser.ID.String(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}

func (service *_IAuthService) signin(userRequest user.ILoginRequest, ctx *context.Context) (user.IAuthUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }
	userResponse, userErr := service.userRepo.Get(userParam, ctx);
	if (userErr != models.IError{}) {
		fmt.Println(userErr)
		return user.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	fmt.Println(userRequest.Password)
	fmt.Println(userResponse.Password)
	passwordErr := utils.CompareHashPassword(userRequest.Password, userResponse.Password);
	if (passwordErr != nil){
		return user.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	accessToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: userResponse.ID.String(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	refreshToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: userResponse.ID.String(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}

	return user.IAuthUserResponse{
		ID: userResponse.ID.String(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}

func (service *_IAuthService) refreshToken(userRequest user.User, ctx *context.Context) (user.IAuthUserResponse, models.IError) {
	accessToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: userRequest.ID.String(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}

	return user.IAuthUserResponse{
		ID: userRequest.ID.String(),
		AccessToken: accessToken,
	}, models.IError{}
}
