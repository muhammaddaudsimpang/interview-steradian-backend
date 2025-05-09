package dto

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type CreateCarReq struct {
	Name      string  `json:"car_name" binding:"required"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required"`
}

type UpdateCarReq struct {
	Name      string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}