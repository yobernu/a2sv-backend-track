package models

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "InProgress"
	Completed  Status = "Completed"
)

//	type Task struct {
//		ID          string `json:"id"`
//		Title       string `json:"title"`
//		Description string `json:"description"`
//		Status      Status `json:"status"`
//	}
type Task struct {
	ID          string  `json:"id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *Status `json:"status,omitempty"`
}
