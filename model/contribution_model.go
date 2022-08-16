package model

type Contribution struct {
	CompanyName   string               `json:"companyName"`
	Contributions []ContributionDetail `json:"contributions"`
}

type ContributionDetail struct {
	Blth         string `json:"blth"`
	Contribution string `json:"contribution"`
}
