package user

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/module"
	"context"
	"strings"

	"gorm.io/gorm"
)

func MigrateModel() string {
	module.DB.AutoMigrate(&CustomSetting{});
	module.DB.AutoMigrate(&Setting{});
	module.DB.AutoMigrate(&Verification{});
	module.DB.AutoMigrate(&SubscriptionPlan{});
	module.DB.AutoMigrate(&Personal{});
	module.DB.AutoMigrate(&Pheripheral{});
	module.DB.AutoMigrate(&User{});
	module.DB.AutoMigrate(&Wallet{});
	
	return "updated user model";
}

type IUserRepository interface {
	Create(request User, ctx *context.Context) (User, models.IError)
	Get(request interface{}, ctx *context.Context) (User, models.IError)
	GetMultiple(search User, option models.IOptions, ctx *context.Context)(IMultipleUser, models.IError)
	Update(user interface{}, update User, ctx *context.Context) (User, models.IError)
}

type PostgresUserRepository struct {
	repo *gorm.DB
}

func NewUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{ 
		repo: module.DB,
	};
}

func (ur *PostgresUserRepository) Create(newUser User, ctx *context.Context) (User, models.IError) {
	err := ur.repo.Create(&newUser).Error;
	if (err != nil) {
		if (strings.Contains(err.Error(), "personals_email_address_key")) {
			return User{}, models.IError{ Field: "email_address", Message: "email address exist already" }
		}
		return User{}, models.IError{ Message: err.Error() }
	}
	return newUser, models.IError{}
}

func (ur *PostgresUserRepository) Get(request interface{}, ctx *context.Context) (User, models.IError) {
	userResponse := User{};
	
	response := ur.repo.Where(&request).
	Preload("Personal").
	Preload("Verification").
	Preload("SubscriptionPlan").
	Preload("Setting").
	Preload("Setting.DefaultOrganisations").
	Preload("Setting.CustomSetting").
	First(&userResponse);
	if (response.Error != nil) {
		return User{}, models.IError{ Message: "unable to get this user" }
	}

	return userResponse, models.IError{};
}

func (ur *PostgresUserRepository) GetMultiple(search User, option models.IOptions, ctx *context.Context) (IMultipleUser, models.IError) {
	var users []User;
	option.Page = (option.Page - 1) * option.Limit;

	userResponse := ur.repo.Where(&search).
	Preload("Setting").
	Preload("Personal").
	Preload("Verification").
	Preload("SubscriptionPlan").
	Limit(int(option.Limit)).Offset(int(option.Page)).Find(&users);
	if (userResponse.Error != nil) {
		return IMultipleUser{}, models.IError{ Message: "unable to get created deals" }
	}

	multipleUserResponse := IMultipleUser{
		TotalUsers: userResponse.RowsAffected,
		Users: users,
		HasNext: ((userResponse.RowsAffected < ((option.Page + 1) * option.Limit))),
	}

	return multipleUserResponse, models.IError{};
}

func (ur *PostgresUserRepository) Update(search interface{}, update User, ctx *context.Context) (User, models.IError) {
	user := User{};
	userResponse := ur.repo.Model(&User{}).Updates(update);
	if (userResponse.Error != nil) {
		return User{}, models.IError{ Message: userResponse.Error.Error() }
	}

	return user, models.IError{};
}