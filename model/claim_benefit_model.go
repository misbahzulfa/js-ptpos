package model

type ClaimBenefit struct {
	BenefitCode                  string `json:"kodeManfaat"`
	ReceiverTypeCode             string `json:"kodeTipePenerima"`
	ProgramCode                  string `json:"kodeProgram"`
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
