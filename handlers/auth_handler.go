package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"book-management/models"
	"book-management/repository"
	"book-management/service"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{
		userService: service.NewUserService(repository.NewUserRepository(db)),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.userService.Register(r.Context(), req.Email, req.Password, req.FullName)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	JSON(w, http.StatusCreated, map[string]any{"id": user.ID, "email": user.Email})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	JSON(w, http.StatusOK, map[string]any{
		"user": user,
	})
}