package repository

import (
	"js-ptpos/entity"
	"js-ptpos/model/request"
	"js-ptpos/model/response"
)

type EmailRepository interface {
	SendEmail(params *request.SendEmailRequest) (*response.SendEmailResponse, error)
	GetDataAfterClaimPTPOS(params *request.GetDataAfterClaimRequest) (*entity.GetDataAfterClaimPTPOSEntity, error)
	GetIDSurvey(params *request.GetDataAfterClaimRequest) (*request.IDSurveyRes, error)
	UpdateTanggalStatusBayar(kodePengajuan string, jenisPengajuan string, userSmile string, noKonfirmasi string) (*response.SendEmailResponse, error)
	CheckExsistKanal40_43(params *request.GetDataAfterClaimRequest) (*response.BeforeSendEmailResponse, error)
	JMOSendNotifPaymentJP(params *request.DataNotifJMORequest) (*response.InsertNotificationResponse, error)
}
