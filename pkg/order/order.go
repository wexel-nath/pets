package order

import "time"

type Order struct {
	ID        int64
	PetID     int64
	Quantity  int64
	ShipDate  time.Time
	Status    string
	Completed bool
}

func (order Order) IsCompleted() bool {
	return order.Status == "delivered"
}
