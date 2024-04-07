package adminUsersRoutes

import (
	"bybu/go-postgres/shared/models"
	"bybu/go-postgres/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IUserManagementController struct {
	service _IUsersService
}

func newUserController() *_IUserManagementController {
	return &_IUserManagementController{ service: *newUserService()}
}

func (controller *_IUserManagementController) getAllUsers(c *fiber.Ctx) error {
		var request models.IOptions;

		ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
		if (errors != nil) {
			return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
		}

		response, err := controller.service.getAllUsers(request, &ctx);
		if (err != models.IError{}) {
			return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err }));
		}

		return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller *_IUserManagementController) banUser(c *fiber.Ctx) error {
	var request user.IPersonalRequest

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.banUser(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller *_IUserManagementController) unbanUser(c *fiber.Ctx) error {
	var request user.IPersonalRequest

	ctx, errors := utils.Validate.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.banUser(request, &ctx);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}
