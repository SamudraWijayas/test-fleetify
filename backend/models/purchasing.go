package models

import "time"

type Purchasing struct {
	ID         uint      `gorm:"primaryKey"`
	Date       time.Time
	SupplierID uint
	UserID     uint
	GrandTotal int
	Details    []PurchasingDetail
}
