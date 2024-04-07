package adminAuthRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/admin"
	"bybu/go-postgres/shared/module/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IAuthController struct {
	service _IAuthService
}

func NewAdminAuthController() *_IAuthController {
	return &_IAuthController{
		service: *NewAdminAuthService(),
	}
}

func (ac _IAuthController) create(c *fiber.Ctx) error {
	var request admin.ICreateRequest;

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(2*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := ac.service.create(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err }));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (ac _IAuthController) login(c *fiber.Ctx) error {
	var request admin.ILoginRequest

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(2*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := ac.service.login(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}