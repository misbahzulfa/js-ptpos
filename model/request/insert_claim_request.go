package request

type InsertClaimRequest struct {
	Email                     string            `json:"email"`
	PhoneNumber               string            `json:"phoneNumber"`
	WorkerCode                string            `json:"workerCode"`
	Kpj                       string            `json:"kpj"`
	IdentityNumber            string            `json:"identityNumber"`
	IdentityType              string            `json:"identityType"`
	ValidIdentity             string            `json:"validIdentity"`
	FullName                  string            `json:"fullName"`
	BirthDate                 string            `json:"birthDate"`
	BirthPlace                string            `json:"birthPlace"`
	Gender                    string            `json:"gender"`
	MotherName                string            `json:"motherName"`
	OfficeCode                string            `json:"officeCode"`
	SegmenCode                string            `json:"segmenCode"`
	CompanyCode               string            `json:"companyCode"`
	DivisionCode              string            `json:"divisionCode"`
	MembershipCode            string            `json:"membershipCode"`
	MembershipDate            string            `json:"membershipDate"`
	ActiveDate                string            `json:"activeDate"`
	NonActiveDate             string            `json:"nonActiveDate"`
	NonActiveCode             string            `json:"nonActiveCode"`
	SumberBlthNonAktif        string            `json:"sumberBlthNonAktif"`
	FlagSubmerNonAktif        string            `json:"flagSubmerNonAktif"`
	TglFlagSumberNonAktif     string            `json:"tglFlagSumberNonAktif"`
	ClaimTypeCode             string            `json:"claimTypeCode"`
	ClaimCauseCode            string            `json:"claimCauseCode"`
	ReportCode                string            `json:"reportCode"`
	ClaimDate                 string            `json:"claimDate"`
	Npwp                      string            `json:"npwp"`
	BankCode                  string            `json:"bankCode"`
	BankName                  string            `json:"bankName"`
	AccountBankNumber         string            `json:"accountBankNumber"`
	AccountBankName           string            `json:"accountBankName"`
	Address                   string            `json:"address"`
	KelurahanCode             string            `json:"kelurahanCode"`
	KecamatanCode             string            `json:"kecamatanCode"`
	KabupatenCode             string            `json:"kabupatenCode"`
	PropinsiCode              string            `json:"propinsiCode"`
	PostalCode                string            `json:"postalCode"`
	ScoreFace                 string            `json:"scoreFace"`
	ScoreFaceLiveness         string            `json:"scoreFaceLiveness"`
	Live                      string            `json:"live"`
	TrxId                     string            `json:"trxId"`
	RefId                     string            `json:"refId"`
	BenefitCode               string            `json:"benefitCode"`
	ReceiverTypeCode          string            `json:"receiverTypeCode"`
	ProgramCode               int               `json:"programCode"`
	TanggalPengembangan       string            `json:"tanggalPengembangan"`
	RatePengembangan          float64           `json:"ratePengembangan"`
	TanggalSaldoAwalTahun     string            `json:"tanggalSaldoAwalTahun"`
	NominalSaldoAwalTahun     float64           `json:"nominalSaldoAwalTahun"`
	NominalSaldoPengembangan  float64           `json:"nominalSaldoPengembangan"`
	NominalSaldoTotal         float64           `json:"nominalSaldoTotal"`
	NominalIuranTahunBerjalan float64           `json:"nominalIuranTahunBerjalan"`
	NominalIuranPengembangan  float64           `json:"nominalIuranPengembangan"`
	NominalIuranTotal         float64           `json:"nominalIuranTotal"`
	NominalSaldoIuranTotal    float64           `json:"nominalSaldoIuranTotal"`
	PersentasePengambilan     float64           `json:"persentasePengambilan"`
	NominalManfaatBisaDiAmbil float64           `json:"nominalManfaatBisaDiAmbil"`
	NominalManfaatDiAmbil     float64           `json:"nominalManfaatDiAmbil"`
	NominalManfaatGross       float64           `json:"nominalManfaatGross"`
	NominalPPH                float64           `json:"nominalPPH"`
	NominalPembulatan         float64           `json:"nominalPembulatan"`
	NominalManfaatNetto       float64           `json:"nominalManfaatNetto"`
	KanalPelayanan            string            `json:"kanalPelayanan"`
	KodeKantorPengajuan       string            `json:"kodeKantorPengajuan"`
	PetugasRekam              string            `json:"petugasRekam"`
	PathUrl                   []DokumenDetilNew `json:"Dokumen"`
	Platform                  string            `json:"platform"`
	DataDetailProbing         []DataProbingNew  `json:"DataProbing"`
	TanggalPengajuan          string            `json:"TanggalPengajuan"`
	KodeBillingPTPOS          string            `json:"KodeBillingPTPOS"`
	ScoreSimilarityNama       string            `json:"ScoreSimilarityNama"`
	JawabanProbing            string            `json:"JawabanProbing"`
	KeteranganPengajuan       string            `json:"KeteranganPengajuan"`
	ChannelId                 string            `json:"ChannelId"`
	KodeKantorPTPOS           string            `json:"KodeKantorPTPOS"`
	PetugasPTPOS              string            `json:"PetugasPTPOS"`
	ChId                      string            `json:"chId"`
	ReqId                     string            `json:"reqId"`
}

type InsertClaimJhtRequest struct {
	SegmenCode          string            `json:"KodeSegmen"`
	Kpj                 string            `json:"KPJ"`
	WorkerCode          string            `json:"WorkerCode"`
	IdentityNumber      string            `json:"NIK"`
	FullName            string            `json:"Nama"`
	BirthDate           string            `json:"TanggalLahir"`
	BirthPlace          string            `json:"TempatLahir"`
	PhoneNumber         string            `json:"NoHP"`
	Address             string            `json:"Alamat"`
	PostalCode          string            `json:"KodePos"`
	KelurahanCode       string            `json:"KodeKelurahan"`
	KecamatanCode       string            `json:"KodeKecamatan"`
	KabupatenCode       string            `json:"KodeKabupaten"`
	PropinsiCode        string            `json:"KodeProvinsi"`
	MotherName          string            `json:"NamaIbuKandung"`
	PathUrl             []DokumenDetilNew `json:"Dokumen"`
	Email               string            `json:"Email"`
	Npwp                string            `json:"NPWP"`
	BankCode            string            `json:"KodeBank"`
	BankName            string            `json:"NamaBank"`
	AccountBankNumber   string            `json:"NomorRekening"`
	AccountBankName     string            `json:"NamaRekening"`
	ScoreFace           string            `json:"ScoreFaceMatch"`
	ScoreFaceLiveness   string            `json:"ScoreFaceLiveness"`
	Live                string            `json:"Live"`
	TrxId               string            `json:"TrxId"`
	RefId               string            `json:"RefId"`
	ClaimTypeCode       string            `json:"ClaimTypeCode"`
	ClaimCauseCode      string            `json:"KodeSebabKlaim"`
	StatusVerklaring    string            `json:"StatusVerklaring"`
	Platform            string            `json:"Platform"`
	DataDetailProbing   []DataProbingNew  `json:"DataProbing"`
	TanggalPengajuan    string            `json:"TanggalPengajuan"`
	KodeBillingPTPOS    string            `json:"KodeBillingPTPOS"`
	ScoreSimilarityNama string            `json:"ScoreSimilarityNama"`
	JawabanProbing      string            `json:"JawabanProbing"`
	KeteranganPengajuan string            `json:"KeteranganPengajuan"`
	ChannelId           string            `json:"channelID"`
	KodeKantorPPTOS     string            `json:"KodeKantorPTPOS"`
	PetugasRekamPTPOS   string            `json:"PetugasPTPOS"`
	ChId                string            `json:"chId"`
	ReqId               string            `json:"reqId"`
}

type DokumenDetilNew struct {
	KodeDokumen    string `json:"KodeDokumen"`
	PathUrlDokumen string `json:"PathUrlDokumen"`
}

type DataProbingNew struct {
	KodeProbing     string `json:"KodeProbing"`
	NomorUrut       string `json:"NomorUrut"`
	ResponseProbing string `json:"ResponseProbing"`
	JawabanProbing  string `json:"JawabanProbing"`
}
