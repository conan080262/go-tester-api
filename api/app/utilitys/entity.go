package utilitys

type Pagination struct {
	TotalRows  int64 `json:"total_rows"`
	PagesTotal int64 `json:"pages_total"`
	PageSkip   int64 `json:"pageskip"`
	PageLimit  int64 `json:"pagelimit"`
	PageIndex  int64 `json:"pageindex"`
}
