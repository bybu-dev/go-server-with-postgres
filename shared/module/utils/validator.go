package utils

import (
	"bybu/go-postgres/shared/models"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type _IValidate struct {
	Body func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors)
	Query func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors)
	Param func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors)
}

var validate = validator.New();



var Validate = _IValidate{
	Body: func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors) {
		ctx, cancel := context.WithTimeout(context.Background(), duration);
		defer cancel();
	
		if err := c.BodyParser(request); err != nil {
			print(err.Error())
			return ctx, models.IErrors{ models.IError{ Message: "unable to decode this body" }};
		}
	
		if err := validate.Struct(request); err != nil {
			var errors = models.IErrors{}
			for _, e:= range err.(validator.ValidationErrors) {
				_error := models.IError{ Field: e.Field(), Message: e.Tag() }
				errors = append(errors, _error)
			}
			return ctx, errors;
		}
	
		return ctx, nil;
	},
	Query: func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors) {
		ctx, cancel := context.WithTimeout(context.Background(), duration);
		defer cancel();
	
		if err := c.BodyParser(request); err != nil {
			print(err.Error())
			return ctx, models.IErrors{ models.IError{ Message: "unable to decode this body" }};
		}
	
		if err := validate.Struct(request); err != nil {
			var errors = models.IErrors{}
			for _, e:= range err.(validator.ValidationErrors) {
				_error := models.IError{ Field: e.Field(), Message: e.Tag() }
				errors = append(errors, _error)
			}
			return ctx, errors;
		}
	
		return ctx, nil;
	},
	Param: func(c *fiber.Ctx, request interface{}, duration time.Duration) (context.Context, models.IErrors) {
		ctx, cancel := context.WithTimeout(context.Background(), duration);
		defer cancel();
	
		if err := c.BodyParser(request); err != nil {
			print(err.Error())
			return ctx, models.IErrors{ models.IError{ Message: "unable to decode this body" }};
		}
	
		if err := validate.Struct(request); err != nil {
			var errors = models.IErrors{}
			for _, e:= range err.(validator.ValidationErrors) {
				_error := models.IError{ Field: e.Field(), Message: e.Tag() }
				errors = append(errors, _error)
			}
			return ctx, errors;
		}
	
		return ctx, nil;
	},
}