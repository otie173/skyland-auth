package postgresql

import (
	"github.com/jmoiron/sqlx"
	"github.com/otie173/skyland-auth/internal/domain/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash)
	return err
}
