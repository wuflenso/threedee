package entity

type PrintRequest struct {
	Id                      int     `json:"id"`
	ItemName                string  `json:"item_name"`
	EstimatedWeight         float32 `json:"estimated_weight"`
	EstimatedFilamentLength float32 `json:"estimated_filament_length"`
	EstimatedDuration       int     `json:"estimated_duration"`
	FileUrl                 string  `json:"file_url"`
	Requestor               string  `json:"requestor"`
	Status                  string  `json:"status"`
}

func NewPrintRequest() *PrintRequest {
	return &PrintRequest{}
}
