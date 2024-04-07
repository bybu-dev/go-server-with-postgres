package admin

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/module"
	"context"
	"strings"

	"gorm.io/gorm"
)


func MigrateModel() string {
	module.DB.AutoMigrate(&Personal{});
	module.DB.AutoMigrate(&Pheripheral{});
	module.DB.AutoMigrate(&Admin{});
	
	return "updated user model";
}

type IAdminRepository interface {
	Create(request ICreateRequest, ctx *context.Context) (Admin, models.IError)
	Get(request interface{}, ctx *context.Context) (Admin, models.IError)
	Update(user interface{}, update Admin, ctx *context.Context) (Admin, models.IError)
}

type AdminMongoRepository struct {
	repo gorm.DB
}

func NewAdminRepository() *AdminMongoRepository {
	return &AdminMongoRepository{
		repo: *module.DB,
	}
}

func (rp AdminMongoRepository) Create(request ICreateRequest, ctx *context.Context) (Admin, models.IError) {
	newAdmin := request.CreateAdmin();

	response := rp.repo.Create(&newAdmin);
	if (response.Error != nil) {
		if (strings.Contains(response.Error.Error(), "E11000")) {
			return Admin{}, models.IError{ Field: "email_address", Message: "admin with this email address already exist" }
		}
		print(response.Error.Error());
		return Admin{}, models.IError{ Message: response.Error.Error() }
	}
	return newAdmin, models.IError{}
}

func (rp AdminMongoRepository) Get(search interface{}, ctx *context.Context) (Admin, models.IError) {
	var admin Admin

	response := rp.repo.Where(search).First(&admin);
	if (response.Error != nil) {
		return Admin{}, models.IError{ Message: response.Error.Error() }
	}

	return admin, models.IError{};
}

func (rp AdminMongoRepository) Update(user interface{}, update Admin, ctx *context.Context) (Admin, models.IError) {
	var admin Admin

	var response = rp.repo.Model(&admin).Updates(update);
	if (response.Error != nil) {
		return Admin{}, models.IError{ Message: response.Error.Error() }
	}

	return admin, models.IError{};
}
