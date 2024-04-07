package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func (user User) ToGeneralResponse() IGeneralUserResponse {
	generalUser := IGeneralUserResponse{
		ID: user.ID,
		Personal: Personal{
			FirstName: user.Personal.FirstName,
			SecondName: user.Personal.SecondName,
			EmailAddress: user.Personal.EmailAddress,
			Username: user.Personal.Username,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return generalUser
}
	
func(user User) ToSecureResponse() ISecureUserResponse {
	secureUser := ISecureUserResponse{
		ID: user.ID,
		Personal: Personal{
			FirstName: user.Personal.FirstName,
			SecondName: user.Personal.SecondName,
			EmailAddress: user.Personal.EmailAddress,
			Username: user.Personal.Username,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Setting: Setting{
			CustomSetting: CustomSetting{
				DefaultTheme: user.Setting.CustomSetting.DefaultTheme,
				IsAcceptingRequest: user.Setting.CustomSetting.IsAcceptingRequest,
			},
		},
		SubscriptionPlan: SubscriptionPlan{
			Status: user.SubscriptionPlan.Status,
			StartDate: user.SubscriptionPlan.StartDate,
			EndDate: user.SubscriptionPlan.EndDate,
		},
	}
	return secureUser
}

func (request IRegisterRequest) CreateUser() User {
	ID, _ := uuid.NewUUID()
	return User{
		ID: ID,
		Personal: Personal{
			FirstName: request.FirstName,
			SecondName: request.SecondName,
			EmailAddress: strings.ToLower(request.EmailAddress),
		},
		Password: request.Password,

		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		
		Pheripheral: Pheripheral{
			IsBanned: false,
			IsVerified: false,
		},
		Setting: Setting{
			CustomSetting: CustomSetting{
				IsAcceptingRequest: false,
			},
		},
		SubscriptionPlan: SubscriptionPlan{
			Status: SubscriptionStatus("BASIC"),
		},
	}
}

func (user IMultipleUser) ToGeneralResponse() IMultipleGeneralResponse {
	var users []IGeneralUserResponse;
	for i := 0; i < len(user.Users); i++ {
		users = append(users, user.Users[i].ToGeneralResponse());
	}
	generalUser := IMultipleGeneralResponse{
		TotalUsers: user.TotalUsers,
		Users: users,
		HasNext: user.HasNext,
	};
	return generalUser
}
	
func(user IMultipleUser) ToSecureResponse() IMultipleSecureResponse {
	var users []ISecureUserResponse;
	for i := 0; i < len(user.Users); i++ {
		users = append(users, user.Users[i].ToSecureResponse());
	}
	secureUser := IMultipleSecureResponse{
		TotalUsers: user.TotalUsers,
		Users: users,
		HasNext: user.HasNext,
	};
	return secureUser
}