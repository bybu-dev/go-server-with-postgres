package adminRoutes

import (
	adminAuthRoutes "bybu/go-postgres/features/admin/auth"
	adminUsersRoutes "bybu/go-postgres/features/admin/users"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(route fiber.Router) {
	adminAuthRoute := route.Group("/auth");
	adminAuthRoutes.UserAuthRoute(adminAuthRoute);

	userManagementRoute := route.Group("/manage/users");
	adminUsersRoutes.AdminUsersRoute(userManagementRoute);
}