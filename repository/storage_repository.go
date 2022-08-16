package repository

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type StorageRepository interface {
	PutObject(params *request.PutObjectRemoteRequest) (*response.PutObjectRemoteResponse, error)
	DeleteObject(params *request.StorageDeleteObjectRemoteRequest) (*response.StorageDeleteObjectRemoteResponse, error)
	UpdateDokumenEC(params *request.UpdateDokumenRequest) (*response.UpdateDokumenResponse, error)
}
