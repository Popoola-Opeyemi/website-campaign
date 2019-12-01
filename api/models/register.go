package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Customers ...
type Customers struct {
	ID           int
	Type         int
	Fullname     string `json:"fullname" validate:"required"`
	OrderType    int    `json:"order_type"`
	PlanType     string `json:"plan_type"`
	Description  string `json:"description"`
	MobileNumber string `json:"mobile_number"`
	EmailAddress string `json:"email" validate:"required"`
	CreatedAt    time.Time
}

// Create a new record
func (s *Customers) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

// GET ...
func (s *Customers) GET(db *gorm.DB, id string) error {
	return db.First(s, "id = ?", id).Error
}

// DELETE ...
func (s *Customers) DELETE(db *gorm.DB) error {
	return db.Delete(s).Error
}

// UPDATE ...
func (s *Customers) UPDATE(db *gorm.DB) error {
	return db.Model(s).Updates(*s).Error
}
