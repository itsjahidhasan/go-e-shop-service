// internal/infrastructure/http/handler/auth_handler.go
package handler

import (
	"encoding/json"
	"go-e-shop-service/internal/usecase"
	"net/http"
)

type AuthHandler struct {
	authUC *usecase.AuthUseCase
}

func NewAuthHandler(authUC *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name, Email, Password string
	}
	json.NewDecoder(r.Body).Decode(&req)
	err := h.authUC.Register(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
