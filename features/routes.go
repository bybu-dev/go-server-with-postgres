package features

import (
	adminRoutes "bybu/go-postgres/features/admin"
	userRoutes "bybu/go-postgres/features/user"

	"github.com/gofiber/fiber/v2"
)

var Routes = func(app *fiber.App) {
	adminRoute := app.Group("/api/admin");
	adminRoutes.AdminRoute(adminRoute)

	userRoute := app.Group("/api/user");
	userRoutes.UserRoute(userRoute)
}