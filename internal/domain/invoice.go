package domain

type Invoice struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"datetime"`
	IdCustomer int     `json:"idcustomer"`
	Total      float64 `json:"total"`
}
