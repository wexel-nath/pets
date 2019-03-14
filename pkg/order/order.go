package order

import "time"

type Order struct {
	ID        int64
	PetID     int64
	Quantity  int32
	ShipDate  time.Time
	Status    string
	Completed bool
}
