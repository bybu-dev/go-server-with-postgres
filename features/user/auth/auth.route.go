package authRoutes

import (
	passwordRoutes "bybu/go-postgres/features/user/auth/password"
	"bybu/go-postgres/shared/module/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	authController := *NewUserAuthController();

	route.Post("/signup", authController.signup);
	route.Post("/signin", authController.signin);
	route.Post("/refreshtoken", middleware.Validate.RefreshRole, authController.refreshToken);

	passwordRoute := route.Group("/password");
	passwordRoutes.UserPasswordRoute(passwordRoute);
}