package charlie

// ParamDelete : parameter of delete charlie
type ParamDelete struct {
	CharlieID uint64 `json:"charlieID" form:"charlieID" query:"charlieID"`
}
