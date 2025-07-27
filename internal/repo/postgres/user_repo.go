package postgres

import (
	"PracticeCrud/internal/domain"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}
func (r *userRepo) FindByUsername(username string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRow("SELECT id, username,password FROM auth_users WHERE username = $1", username).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, err
}
