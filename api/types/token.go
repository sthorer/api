package types

import "github.com/sthorer/api/ent"

type NewTokenRequest struct {
	Name string `json:"name" validate:"required"`
}

type TokenSecretResponse struct {
	*ent.Token
	Secret string `json:"secret"`
}
