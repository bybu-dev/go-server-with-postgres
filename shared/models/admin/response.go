package admin

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

	Personal Personal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IGeneralUserResponse struct {
	ID uuid.UUID `json:"id"`

	Personal Personal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}