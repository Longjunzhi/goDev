package services

import (
	"context"
	"net/http"
)

type LoginByPasswordRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginByPasswordResponse struct {
	Token string `json:"token"`
}

func LoginByPassword(ctx context.Context, req *LoginByPasswordRequest) (resp *LoginByPasswordResponse, code int, err error) {
	resp.Token = "123456"
	return resp, http.StatusOK, nil
}
