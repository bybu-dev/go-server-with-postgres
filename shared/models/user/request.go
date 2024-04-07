package user

import (
	"time"
)

type IRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required,email"`
    UserRoles []UserRole `json:"user_roles"`
	Password string `json:"password" validate:"required"`
}

type ILoginRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ICreateTokenRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type ISendResetCodeRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type IVerifyResetCodeRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Code string `json:"auth_code" validate:"required"`
	Password string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type IPersonalRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type IUpdatePersonal struct {
	FirstName string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required,email"`
	ProfileImage string `json:"profile_image"`
	Username string `json:"username"`
}

type IUserCodeInjection struct {
	AuthenticationCode string `json:"authentication_code" bson:"pheripheral.authentication_code"`
	Timeout time.Time `json:"timeout" bson:"pheripheral.timeout"`
	IsBanned bool `json:"is_banned" bson:"is_banned,false"`
}