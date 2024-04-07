package passwordRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"context"
)

type _IPasswordService struct {
	userRepo user.IUserRepository
}

func NewUserPasswordService() *_IPasswordService {
	return &_IPasswordService{ userRepo: user.NewUserRepository() }
}

func (service _IPasswordService) sendResetCode(request user.ISendResetCodeRequest, ctx *context.Context) (user.IPasswordAuthResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: request.EmailAddress };
	updatedParam := user.User{ 
		// Pheripheral: &user.Pheripheral{AuthenticationCode: "0000", Timeout: time.Now()},
	};

	var userResponse, err = service.userRepo.Update( userParam, updatedParam, ctx);
	if (err != models.IError{}) {
		return user.IPasswordAuthResponse{}, models.IError{ 
			Field: "email_address",
			Message: "email address was not found",
		}
	}

	mailErr := utils.EmailSender.SendVerification("0000");
	if (mailErr != nil) {
		return user.IPasswordAuthResponse{}, models.IError{ 
			Field: "email_address",
			Message: "unable to send verification code",
		}
	}

	return user.IPasswordAuthResponse{
		ID: userResponse.ID.String(),
		EmailAddress: request.EmailAddress,
		Message: "code has been sent to this email",
	}, models.IError{}
}

func (service _IPasswordService) resetPassword(userRequest user.IVerifyResetCodeRequest, ctx *context.Context) (user.IPasswordAuthResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	var userResponse user.User
	userResponse, responseErr := service.userRepo.Get(&userParam, ctx);
	if (responseErr != models.IError{}) {
		return user.IPasswordAuthResponse{}, models.IError{
			Field: "email_address",
			Message: "email is not registered with us",
		}
	}

	// if (userResponse.Pheripheral.AuthenticationCode != userRequest.Code) {
	// 	return user.IPasswordAuthResponse{}, models.IError{ Field: "email_address", Message: "invalid code"}
	// }
	// if (userResponse.Pheripheral.Timeout.Before(time.Now().Add(time.Hour*24*7))) {
	// 	return user.IPasswordAuthResponse{}, models.IError{ Field: "email_address", Message: "code expired"}
	// }

	return user.IPasswordAuthResponse{
		ID: userResponse.ID.String(),
		EmailAddress: userRequest.EmailAddress,
		Message: "code has been sent to this email",
	}, models.IError{}
}
