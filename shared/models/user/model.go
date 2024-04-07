package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string;
const (
    DEVELOPER UserRole = "DEVELOPER";
    CTO UserRole = "CTO";
    CEO UserRole = "CEO";
    MARKETER UserRole = "MARKETER";
    PROJECT_MANAGER UserRole = "PROJECT MANAGER";
);

type SubscriptionStatus string;
const (
    PRO SubscriptionStatus = "PRO";
    BASIC SubscriptionStatus = "BASIC";
    ENTERPRICE SubscriptionStatus = "ENTERPRICE";
)

type CustomSetting struct {
	gorm.Model

	UserID uuid.UUID `gorm:"foreignKey:UserID"`
    DefaultTheme string `json:"default_theme"`;
    IsAcceptingRequest bool `json:"is_accepting_request"`;

	UpdatedAt time.Time `json:"-"`
    DeletedAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type Setting struct {
	gorm.Model

	UserID uuid.UUID
    CustomSetting CustomSetting `json:"custom_setting" gorm:"foreignKey:UserID"`

	UpdatedAt time.Time `json:"-"`
    DeletedAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type Verification struct {
	// gorm.Model

	UserID uuid.UUID
    Code string `json:"code"`;
    Timeout time.Duration `json:"timeout"`;

	UpdatedAt time.Time `json:"-"`
    DeletedAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type SubscriptionPlan struct {
	gorm.Model

	UserID uuid.UUID
    Status SubscriptionStatus `json:"status" gorm:"type:text"`;
    StartDate time.Time `json:"start_date"`;
    EndDate time.Time `json:"end_date"`;

	UpdatedAt time.Time `json:"-"`
    DeletedAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type Personal struct {
	gorm.Model

	UserID uuid.UUID
	FirstName string `json:"first_name" validate:"required"`
	SecondName string `json:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required,email" gorm:"unique"`
	ProfileImage string `json:"profile_image"`
	Username string `json:"username"`

	UpdatedAt time.Time `json:"-"`
    DeletedAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type Pheripheral struct {
	gorm.Model
	
	UserID uuid.UUID
	AuthenticationCode string `json:"authentication_code"`
	Timeout time.Time `json:"timeout"`
	IsBanned bool `json:"is_banned,false"`
    IsVerified bool `json:"is_verified"`

    UpdatedAt time.Time `json:"-"`
    DeleteAt time.Time `json:"-"`
    CreatedAt time.Time `json:"-"`
}

type User struct {
	gorm.Model

	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
    Personal Personal `json:"personal"`;
	Password string `json:"password"`;

    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"delete_at"`
    CreatedAt time.Time `json:"created_at"`
	
    Verification Verification `json:"verification"`
	Pheripheral Pheripheral `json:"pheripheral"`
    Setting Setting `json:"setting"`
    SubscriptionPlan SubscriptionPlan `json:"subscription_plan"`
}

type Wallet struct {
	gorm.Model
	
	UserID uuid.UUID
	Balance int `json:"balance"`
	TotalBalance int `json:"total_balance"`
}

type IMultipleUser struct {
	TotalUsers int64 `json:"total_users"`
	Users []User `json:"users"`
	HasNext bool `json:"has_next"`
}


