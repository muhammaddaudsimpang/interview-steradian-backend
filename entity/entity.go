package entity

import "time"

var (
	MsgDataOK = "data successfully retrieved"
	MsgCarCreated = "car successfully created"
	MsgCarUpdated = "car successfully updated"
	MsgCarDeleted = "car successfully deleted"
)

type Car struct {
	ID        int     `json:"id"`
	Name      string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}

type Order struct {
	ID              int       `json:"id"`
	CarID           int       `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}

type CarWithOrder struct {
	Car   Car   `json:"car"`
	Order Order `json:"order"`
}
