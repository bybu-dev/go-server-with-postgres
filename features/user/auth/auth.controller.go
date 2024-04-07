package authRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IAuthController struct {
	service _IAuthService
}

func NewUserAuthController() *_IAuthController {
	return &_IAuthController{
		service: *NewUserAuthService(),
	}
}

func (controller _IAuthController) signup(c *fiber.Ctx) error {
	var request user.IRegisterRequest;

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.signup(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IAuthController) signin(c *fiber.Ctx) error {
	var request user.ILoginRequest

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.signin(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IAuthController) refreshToken(c *fiber.Ctx) error {
	request, _ := c.Locals("user").(user.User);
	
	ctx, ctxErr := context.WithTimeout(context.Background(), time.Duration(10*time.Second));
	defer ctxErr();

	response, err := controller.service.refreshToken(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}