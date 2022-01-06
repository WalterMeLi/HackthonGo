package domain

type Sale struct {
	ID        int     `json:"id"`
	IdInvoice int     `json:"idinvoice"`
	IdProduct int     `json:"idproduct"`
	Quantity  float64 `json:"quantity"`
}
