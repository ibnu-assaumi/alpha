package charlie

// ParamGet : parameter of get charlie
type ParamGet struct {
	CharlieID   uint64 `json:"charlieID" query:"charlieID"`
	CharlieName string `json:"charlieName" query:"charlieName"`
	Page        int    `json:"page" query:"page"`
	Limit       int    `json:"limit" query:"limit"`
	Descending  bool   `json:"descending" query:"descending"`
	OrderBy     string `json:"orderBy" query:"orderBy"`
}
