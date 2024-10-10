package types

type ExportExcelCol struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ExportFileItem struct {
	Value  string `json:"value"`
	Rename string `json:"rename"`
}

type ExportExcelRequest struct {
	UserId       uint32              `json:"user_id"`
	DepartmentId uint32              `json:"department_id"`
	Scene        string              `json:"scene"`
	Name         string              `json:"name"`
	Files        []*ExportFileItem   `json:"files"`
	Headers      []string            `json:"headers"`
	Rows         [][]*ExportExcelCol `json:"rows"`
}
