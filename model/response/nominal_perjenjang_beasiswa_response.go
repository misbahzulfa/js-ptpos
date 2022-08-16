package response

type NominalPerjenjangBeasiswaResponse struct {
	StatusCode                string                      `json:"statusCode"`
	StatusDesc                string                      `json:"statusDesc"`
	NominalPerjenjangBeasiswa []NominalPerjenjangBeasiswa `json:"dataNominalPerjenjangBeasiswa"`
}

type NominalPerjenjangBeasiswa struct {
	NominalManfaat string `json:"nominalManfaat"`
	Keterangan     string `json:"keterangan"`
}
