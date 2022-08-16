package entity

type ClaimBenefitDetailEntity struct {
	BenefitCode                  string `json:"benefitCode"`
	ReceiverTypeCode             string `json:"receiverTypeCode"`
	ProgramCode                  string `json:"programCode"`
	ReportCode                   string `json:"reportCode"`
	TanggalPengembangan          string `json:"tanggalPengembangan"`
	RatePengembangan             string `json:"ratePengembangan"`
	TanggalSaldoAwalTahun        string `json:"tanggalSaldoAwalTahun"`
	NominalSaldoAwalTahun        string `json:"nominalSaldoAwalTahun"`
	NominalSaldoPengembangan     string `json:"nominalSaldoPengembangan"`
	NominalSaldoTotal            string `json:"nominalSaldoTotal"`
	NominalIuranTahunBerjalan    string `json:"nominalIuranTahunBerjalan"`
	NominalIuranPengembangan     string `json:"nominalIuranPengembangan"`
	NominalIuranTotal            string `json:"nominalIuranTotal"`
	NominalSaldoIuranTotal       string `json:"nominalSaldoIuranTotal"`
	PersentasePengambilan        string `json:"persentasePengambilan"`
	NominalManfaatMaxBisaDiAmbil string `json:"nominalManfaatMaxBisaDiAmbil"`
	NominalManfaatDiAmbil        string `json:"nominalManfaatDiAmbil"`
	NominalManfaatGross          string `json:"nominalManfaatGross"`
	NominalPPH                   string `json:"nominalPPH"`
	NominalPembulatan            string `json:"nominalPembulatan"`
	NominalManfaatNetto          string `json:"nominalManfaatNetto"`
}
