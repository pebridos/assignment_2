package models

import "time"

type Orders struct {
	OrderID      uint      `json:"orderId" gorm:"primaryKey"`
	CostumerName string    `json:"costumerName" gorm:"not null;type:varchar(255)"`
	OrderedAt    time.Time `json:"orderAt"`
	Items        []Items   `json:"items"`
}
