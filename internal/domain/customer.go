package domain

type Customer struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name`
	FirstName string `json:"first_name"`
	Condition string `json:"condition"`
}
