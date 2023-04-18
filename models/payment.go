package models

import "time"

type Payment struct {
	ID          int       `json:"id"`
	Customer_Id int       `json:"customer_id"`
	Bill        float32   `json:"bill"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
