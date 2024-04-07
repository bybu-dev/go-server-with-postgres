package profileRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IProfileController struct {
	service _IProfileService
}

func NewProfileController() *_IProfileController {
	return &_IProfileController{ service: *NewUserProfileService() }
}

func (controller _IProfileController) getUser(c *fiber.Ctx) error {
	ctx, ctxErr := context.WithTimeout(context.Background(), time.Second);
	defer ctxErr();

	response, errResponse := controller.service.getProfile(c.Locals("user").(user.User), &ctx);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}

func (controller _IProfileController) updateUser(c *fiber.Ctx) error {
	var request user.Personal;
	ctx, err:= utils.Validate.Body(c, request, time.Duration(10*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err))
	}

	response, errResponse := controller.service.updateProfile(c.Locals("user").(user.User), request, &ctx);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}