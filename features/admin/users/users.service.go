package adminUsersRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"context"
)

type _IUsersService struct {
	userRepo user.IUserRepository
}

func newUserService() *_IUsersService {
	return &_IUsersService{
		userRepo: user.NewUserRepository(),
	}
}

func (service *_IUsersService) getAllUsers(request models.IOptions, ctx *context.Context) (user.IMultipleGeneralResponse, models.IError) {
	multipleUsers, userErr := service.userRepo.GetMultiple(user.User{}, request, ctx);
	if (userErr != models.IError{}) {
		return user.IMultipleGeneralResponse{}, userErr
	}

	return multipleUsers.ToGeneralResponse(), models.IError{}
}

func (service *_IUsersService) banUser(userRequest user.IPersonalRequest, ctx *context.Context) (user.IGeneralUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	userResponse, userErr := service.userRepo.Update(userParam, user.User{
		// Pheripheral: &user.Pheripheral{ IsBanned: true },
	}, ctx);

	if (userErr != models.IError{}) {
		return user.IGeneralUserResponse{}, models.IError{ Message: "unable to ban this user" }
	}

	return userResponse.ToGeneralResponse(), models.IError{}
}

func (service *_IUsersService) unbanUser(userRequest user.IPersonalRequest, ctx *context.Context) (user.IGeneralUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	userResponse, userErr := service.userRepo.Update(userParam, user.User{
		// Pheripheral: &user.Pheripheral{ IsBanned: true },
	}, ctx);

	if (userErr != models.IError{}) {
		return user.IGeneralUserResponse{}, models.IError{ Message: "unable to ban this user" }
	}

	return userResponse.ToGeneralResponse(), models.IError{}
}