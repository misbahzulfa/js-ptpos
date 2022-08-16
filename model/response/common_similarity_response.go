package response

type CommonSimilarityResponse struct {
	Ret      string  `json:"ret"`
	MinScore float64 `json:"minScore"`
	Score    float64 `json:"score"`
	Nama2    string  `json:"nama2"`
	Nama1    string  `json:"nama1"`
	Match    bool    `json:"match"`
	Msg      string  `json:"msg"`
}
