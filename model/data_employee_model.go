package model

type DataEmployee struct {
	AsikCode                         string `json:"asikCode"`
	IdentityNumber                   string `json:"identityNumber"`
	IdentityType                     string `json:"identityType"`
	ValidIdentity                    string `json:"validIdentity"`
	FullName                         string `json:"fullName"`
	Gender                           string `json:"gender"`
	BirthPlace                       string `json:"birthPlace"`
	BirthDate                        string `json:"birthDate"`
	MotherName                       string `json:"motherName"`
	MaritalStatus                    string `json:"maritalStatus"`
	Address                          string `json:"address"`
	KelurahanCode                    string `json:"kelurahanCode"`
	KecamatanCode                    string `json:"kecamatanCode"`
	KabupatenCode                    string `json:"kabupatenCode"`
	PropinsiCode                     string `json:"propinsiCode"`
	PostalCode                       string `json:"postalCode"`
	PhoneNumber                      string `json:"phoneNumber"`
	Email                            string `json:"email"`
	Npwp                             string `json:"npwp"`
	PassportNumber                   string `json:"passportNumber"`
	PassportExpired                  string `json:"passportExpired"`
	BankCode                         string `json:"bankCode"`
	BankName                         string `json:"bankName"`
	AccountBankNumber                string `json:"accountBankNumber"`
	AccountBankName                  string `json:"accountBankName"`
	LastEducationCode                string `json:"lastEducationCode"`
	ReligionCode                     string `json:"religionCode"`
	BloodGroup                       string `json:"bloodGroup"`
	EmergencyContactName             string `json:"emergencyContactName"`
	EmergencyContactPhoneNumber      string `json:"emergencyContactPhoneNumber"`
	EmergencyContactAddress          string `json:"emergencyContactAddress"`
	EmergencyContactRelationshipCode string `json:"emergencyContactRelationshipCode"`
	EmergencyContactKelurahanCode    string `json:"emergencyContactKelurahanCode"`
	EmergencyContactKecamatanCode    string `json:"emergencyContactKecamatanCode"`
	EmergencyContactKabupatenCode    string `json:"emergencyContactKabupatenCode"`
	EmergencyContactPropinsiCode     string `json:"emergencyContactPropinsiCode"`
	EmergencyContactPostalCode       string `json:"emergencyContactPostalCode"`
	WorkerCode                       string `json:"workerCode"`
	Kpj                              string `json:"kpj"`
	SegmenCode                       string `json:"segmenCode"`
	DivisionCode                     string `json:"divisionCode"`
	CompanyCode                      string `json:"companyCode"`
	MembershipCode                   string `json:"membershipCode"`
	OfficeCode                       string `json:"officeCode"`
	Npp                              string `json:"npp"`
	MembershipDate                   string `json:"membershipDate"`
	ActiveDate                       string `json:"activeDate"`
	NonActiveDate                    string `json:"nonActiveDate"`
	NonActiveCode                    string `json:"nonActiveCode"`
	FlagSipp                         string `json:"flagSipp"`
	Active                           string `json:"active"`
	JhtBalance                       string `json:"jhtBalance"`
}
