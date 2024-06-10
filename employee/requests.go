package employee

type CreateEmployeeRequest struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}