package passwordRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IPasswordController struct {
	service _IPasswordService
}

func NewUserPasswordController() *_IPasswordController {
	return &_IPasswordController{
		service: *NewUserPasswordService(),
	}
}

func (controller _IPasswordController) sendResetCode(c *fiber.Ctx) error {
	var request user.ISendResetCodeRequest
	
	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.sendResetCode(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IPasswordController) resetPassword(c *fiber.Ctx) error {
	var request user.IVerifyResetCodeRequest

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.resetPassword(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}