package charlie

// ParamUpdate : parameter of update existing charlie record
type ParamUpdate struct {
	CharlieID   uint64 `json:"charlieID" form:"charlieID" query:"charlieID"`
	CharlieName string `json:"charlieName" form:"charlieName" query:"charlieName"`
}
