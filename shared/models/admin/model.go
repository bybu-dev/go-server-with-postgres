package admin

import (
	"time"

	"github.com/google/uuid"
)

type Personal struct {
	UserID uuid.UUID
	FirstName string `json:"first_name" bson:"first_name" validate:"required"`
	SecondName string `json:"second_name" bson:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" bson:"email_address,unique" validate:"required,email"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
	Username string `json:"username" bson:"username"`
}

type Pheripheral struct {
	UserID uuid.UUID
	AuthenticationCode string `json:"authentication_code" bson:"authentication_code"`
	Timeout time.Time `json:"timeout" bson:"timeout"`
}

type Admin struct {
	ID uuid.UUID `json:"id" bson:"_id,omitempty"`

	Personal Personal `json:"personal" bson:"personal,omitempty"`
	Pheripheral Pheripheral `json:"pheripheral" bson:"pheripheral,omitempty"`

	Password string `json:"password" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}