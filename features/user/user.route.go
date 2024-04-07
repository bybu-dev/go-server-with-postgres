package userRoutes

import (
	authRoutes "bybu/go-postgres/features/user/auth"
	personalRoutes "bybu/go-postgres/features/user/personal"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	authRoute := route.Group("/auth");
	authRoutes.Routes(authRoute);

	personalRoute := route.Group("/personal");
	personalRoutes.Route(personalRoute);
}