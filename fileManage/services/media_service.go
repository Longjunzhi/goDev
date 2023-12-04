package services

import (
	"Img/model"
	"context"
	"mime/multipart"
	"net/http"
)

type MediaRequest struct {
	Files []*multipart.FileHeader `json:"files"`
}

type MediaResponse struct {
	Media []model.Media `json:"media"`
}

func StoreMedia(ctx context.Context, req *MediaRequest) (resp *LoginByPasswordResponse, code int, err error) {
	return resp, http.StatusOK, nil
}
