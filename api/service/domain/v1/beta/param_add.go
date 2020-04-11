package beta

type ParamAdd struct {
	BetaName      string   `json:"betaName"`
	CharlieIDList []uint64 `json:"charlieIDList"`
}
