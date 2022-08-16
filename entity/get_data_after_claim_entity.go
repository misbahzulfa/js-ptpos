package entity

type GetDataAfterClaimEntity struct {
	SubmissionCode    string `json:"submissionCode"`
	ClaimCode         string `json:"claimCode"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	Kpj               string `json:"kpj"`
	IdentityNumber    string `json:"identityNumber"`
	Fullname          string `json:"fullName"`
	Birthdate         string `json:"birthdate"`
	Birthplace        string `json:"birthplace"`
	Gender            string `json:"gender"`
	Npwp              string `json:"npwp"`
	BankCode          string `json:"bankCode"`
	BankName          string `json:"bankName"`
	AccountBankNumber string `json:"accountBankNumber"`
	AccountBankName   string `json:"accountBankName"`
	TotalTransfer     string `json:"totalTransfer"`
	Address           string `json:"address"`
	PaymentStatus     string `json:"paymentStatus"`
	PaymentDate       string `json:"paymentDate"`
	ActiveDate        string `json:"activeDate"`
	NonActiveDate     string `json:"nonActiveDate"`
	MembershipDate    string `json:"membershipDate"`
}

type GetDataAfterClaimPTPOSEntity struct {
	KodePengajuan   string            `json:"kodePengajuan"`
	TipePengajuan   string            `json:"tipePengajuan"`
	KodeKlaim       string            `json:"kodeKlaim"`
	KPJ             string            `json:"kpj"`
	NamaTK          string            `json:"namaTK"`
	WaktuPembayaran string            `json:"waktuPembayaran"`
	BLTHPengajuan   string            `json:"blthPengajuan"`
	DataPenerima    []GetDataPenerima `json:"dataPenerima"`
	DataEmail       []GetEmailData    `json:"dataEmail"`
}

type GetDataPenerima struct {
	NamaPenerima     string `json:"namaPenerima"`
	NamaBank         string `json:"bankTujuan"`
	NoRekening       string `json:"rekeningTujuan"`
	NamaRekening     string `json:"rekeningPenerima"`
	JumlahPembayaran string `json:"jumlahPembayaran"`
}

type GetEmailData struct {
	EmailSubject string `json:"emailSubject"`
	EmailContent string `json:"emailContent"`
	Email        string `json:"email"`
}
