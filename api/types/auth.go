package types

import "github.com/sthorer/api/ent"

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthResponse struct {
	*ent.User
	Token string `json:"token"`
}
