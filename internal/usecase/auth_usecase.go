// internal/usecase/auth_usecase.go
package usecase

import (
	"context"
	"errors"
	"go-e-shop-service/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo domain.UserRepository
}

func NewAuthUseCase(repo domain.UserRepository) *AuthUseCase {
	return &AuthUseCase{repo}
}

func (uc *AuthUseCase) Register(ctx context.Context, name, email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{Name: name, Email: email, Password: string(hashed)}
	return uc.repo.Create(ctx, user)
}

func (uc *AuthUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.repo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
