package entity

type PrintRequest struct {
	Id                      int
	ItemName                string
	EstimatedWeight         float32
	EstimatedFilamentLength float32
	EstimatedDuration       int
	FileUrl                 string
	Requestor               string
	Status                  string
}
