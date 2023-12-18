package services

import (
	"Img/databases"
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

type MediaGetRequest struct {
	Count    int `json:"count"`
	Offset   int `json:"offset"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type MediaGetResponse struct {
	Media []model.Media `json:"media"`
}

func MediaGet(ctx context.Context, req *MediaGetRequest) (resp *MediaGetResponse, code int, err error) {
	resp = &MediaGetResponse{}
	medias, _, err := databases.MediaGetByOffset(context.Background(), req.Offset, req.Count)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}
	resp.Media = medias
	return resp, http.StatusOK, nil
}
