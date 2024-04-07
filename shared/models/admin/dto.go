package admin

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func (user Admin) ToAdminResponse() ISecureUserResponse {
		secureUser := ISecureUserResponse{
			ID: user.ID,
			Personal: user.Personal,

			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		return secureUser
}

func (request ICreateRequest) CreateAdmin() Admin {
	ID, _ := uuid.NewUUID();
	return Admin{
		ID: ID,
		Personal: Personal{
			FirstName: request.Personal.FirstName,
			SecondName: request.Personal.SecondName,
			EmailAddress: strings.ToLower(request.Personal.EmailAddress),
		},
		Password: request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}