package walletRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IProfileController struct {
	service _IProfileService
}

func NewUserWalletController() *_IProfileController {
	return &_IProfileController{ service: *NewUserWalletService() }
}

func (controller _IProfileController) fund(c *fiber.Ctx) error {
	var request user.Personal;
	ctx, err:= utils.Validate.Body(c, request, time.Duration(2*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err))
	}

	response, errResponse := controller.service.fund(c.Locals("user").(user.User), request, &ctx);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}

func (controller _IProfileController) withdraw(c *fiber.Ctx) error {
	var request user.Personal;
	ctx, err:= utils.Validate.Body(c, request, time.Duration(2*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err))
	}

	response, errResponse := controller.service.withdraw(c.Locals("user").(user.User), request, &ctx);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}