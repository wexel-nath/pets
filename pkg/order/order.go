package order

import "time"

type Order struct {
	ID        int64     `json:"id"`
	PetID     int64     `json:"petId"`
	Quantity  int64     `json:"quantity"`
	ShipDate  time.Time `json:"shipDate"`
	Status    string    `json:"status"`
	Completed bool      `json:"complete"`
}

// Derive completed from status
func (order Order) IsCompleted() bool {
	return order.Status == "delivered"
}
