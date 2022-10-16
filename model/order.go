package model

type Order struct {
	Order_ID      int    `gorm:"primaryKey" json:"order_id"`
	Customer_Name string `gorm:"not null;type:varchar(50)" json:"customer_name"`
	Ordered_at    string `gorm:"not null;type:varchar(50)" json:"ordered_at"`
}
