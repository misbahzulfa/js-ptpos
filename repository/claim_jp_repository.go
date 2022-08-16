package repository

import (
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type ClaimJPRepository interface {
	JPCheckEligible(params *request.CheckEligibleJPBerkalaRequest) (*response.CheckEligibleJPBerkalaResponse, error)
	JPCheckJumlahKlaimBerkalaByNIKPelapor(params *request.CheckJumlahKlaimJPBerkalaNikPelaporRequest) (*response.CheckJumlahKlaimJPBerkalaNikPelaporResponse, error)
	InsertConfirmationJPBerkala(params *request.InsertJPConfirmationRequest) (*response.InsertKonfirmasiJPResponse, error)
	CheckStatusKonfirmasiExsist(workerCode string) int
}
