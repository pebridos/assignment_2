package models

type Items struct {
	ItemID          uint `json:"LineItemId" gorm:"primaryKey"`
	OrdersID        uint
	ItemCode        string `json:"itemCode" gorm:"not null;type:varchar(10)"`
	ItemDescription string `json:"description" gorm:"not null;type:varchar(255)"`
	ItemQuantity    int8   `json:"quantity" gorm:"not null;type:integer"`
}
