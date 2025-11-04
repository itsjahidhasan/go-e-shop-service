// internal/repository/user_repository.go
package repository

import (
	"context"
	"database/sql"
	"go-e-shop-service/internal/domain"
)

type UserRepoPostgres struct {
	db *sql.DB
}

func NewUserRepoPostgres(db *sql.DB) *UserRepoPostgres {
	return &UserRepoPostgres{db}
}

func (r *UserRepoPostgres) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password, created_at)
			  VALUES ($1, $2, $3, NOW())`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepoPostgres) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password, created_at FROM users WHERE email=$1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
