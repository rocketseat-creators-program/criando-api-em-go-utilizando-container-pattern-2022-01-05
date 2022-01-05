package structs

type Installment struct {
	ID     string  `json:"id"`
	Value  float64 `json:"value"`
	DueDay int     `json:"due_day"`
}
