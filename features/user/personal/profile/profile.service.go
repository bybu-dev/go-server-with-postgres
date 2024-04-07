package profileRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"context"
)

type _IProfileService struct {
	userRepo user.IUserRepository
}

func NewUserProfileService() *_IProfileService {
	return &_IProfileService{ userRepo: user.NewUserRepository() }
}

func (service *_IProfileService) getProfile(me user.User, ctx *context.Context) (user.ISecureUserResponse, models.IError) {
	return me.ToSecureResponse(), models.IError{};
}

func (service *_IProfileService) updateProfile(me user.User, personal user.Personal, ctx *context.Context) (user.ISecureUserResponse, models.IError) {
	request := user.User{ ID: me.ID };
	updateRequest := user.User{ Personal: personal };
	
	updatedUser, err := service.userRepo.Update(request, updateRequest, ctx);
	if (err != models.IError{}) {
		return user.ISecureUserResponse{}, err;
	}

	return updatedUser.ToSecureResponse(), models.IError{};
}