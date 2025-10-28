package repositories

import "github.com/otie173/skyland-auth/internal/domain/models"

type UserRepository interface {
	Create(user models.User) error
	//FindByUsername(username string) (models.User, error)
	//FindByEmail(email string) (models.User, error)
}
