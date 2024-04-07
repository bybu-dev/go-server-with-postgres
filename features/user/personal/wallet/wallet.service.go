package walletRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"context"
)

type _IProfileService struct {
	userRepo user.IUserRepository
}

func NewUserWalletService() *_IProfileService {
	return &_IProfileService{ userRepo: user.NewUserRepository() }
}

func (service *_IProfileService) fund(me user.User, personal user.Personal, ctx *context.Context) (user.User, models.IError) {
	request := user.User{ ID: me.ID };
	updateRequest := user.User{ Personal: personal };
	
	updatedUser, err := service.userRepo.Update(request, updateRequest, ctx);
	if (err != models.IError{}) {
		return user.User{}, err;
	}

	return updatedUser, models.IError{};
}

func (service *_IProfileService) withdraw(me user.User, personal user.Personal, ctx *context.Context) (user.User, models.IError) {
	request := user.User{ ID: me.ID };
	updateRequest := user.User{ Personal: personal };
	
	updatedUser, err := service.userRepo.Update(request, updateRequest, ctx);
	if (err != models.IError{}) {
		return user.User{}, err;
	}

	return updatedUser, models.IError{};
}