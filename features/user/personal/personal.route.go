package personalRoutes

import (
	profileRoutes "bybu/go-postgres/features/user/personal/profile"
	walletRoutes "bybu/go-postgres/features/user/personal/wallet"

	"github.com/gofiber/fiber/v2"
)

func Route(route fiber.Router) {
	profileRoute := route.Group("/profile");
	profileRoutes.Route(profileRoute);

	walletRoute := route.Group("/wallet");
	walletRoutes.Route(walletRoute);
}