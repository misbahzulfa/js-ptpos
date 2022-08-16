package model

type CheckAccountBank struct {
	IsSuccessful bool   `json:"isSuccessful"`
	Message      string `json:"message"`
	Signature    string `json:"signature"`
}
