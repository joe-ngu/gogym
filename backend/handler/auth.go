package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joe-ngu/gogym/store"
	"github.com/joe-ngu/gogym/types"
)

type AuthHandler struct {
	db store.DB
}

func NewAuthHandler(db store.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling LOGIN request - Method:", r.Method)
	var req types.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	user, err := h.db.GetUserByUsername(req.UserName)
	if err != nil {
		return err
	}

	if !user.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := CreateJWT(user.ID, user.UserName)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
		UserName: user.UserName,
		Token:    token,
	}
	return writeJSON(w, http.StatusOK, resp)
}
