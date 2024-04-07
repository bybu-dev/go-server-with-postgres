package user

import (
	"time"

	"github.com/google/uuid"
)

type IAuthUserResponse struct {
	ID string `json:"id"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type IPasswordAuthResponse struct {
	ID string `json:"id"`
	EmailAddress string `json:"account_type"`
	Message string `json:"message"`
}

type ISecureUserResponse struct {
	ID uuid.UUID `json:"id"`
    UserRoles []UserRole `json:"user_roles"`
    Personal Personal `json:"personal"`;

    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
	
    Verification Verification `json:"verification"`
    Setting Setting `json:"setting"`
    SubscriptionPlan SubscriptionPlan `json:"subscription_plan"`
}

type IGeneralUserResponse struct {
	ID uuid.UUID `json:"id"`

	Personal Personal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IMultipleSecureResponse struct {
	TotalUsers int64 `json:"total_users"`
	Users []ISecureUserResponse `json:"users"`
	HasNext bool `json:"has_next"`
}

type IMultipleGeneralResponse struct {
	TotalUsers int64 `json:"total_users"`
	Users []IGeneralUserResponse `json:"users"`
	HasNext bool `json:"has_next"`
}