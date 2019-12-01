package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Roles with lowest value have highest privilege
const (
	// RoleSuperAdmin for super admin
	RoleSuperAdmin int64 = iota

	// RoleAdmin for admin
	RoleAdmin

	// RoleManager for manager
	RoleManager

	// RoleStaff for staff
	RoleStaff
)

// Authentications ...
type Authentications struct {
	ID          int       `json:"id"`
	RoleID      int64     `json:"role_id"`
	Email       string    `json:"username"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"created_at"`
}

// LoginForm ...
type LoginForm struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// All ... gets all new record
func (c *Authentications) All(db *gorm.DB) (admin []*Authentications, err error) {
	err = db.Find(&admin).Error
	return
}

// GET ...
func (c *Authentications) GET(db *gorm.DB, id string) error {
	return db.First(c, "id = ?", id).Error
}

// Create a new record
func (c *Authentications) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

// GetByEmail ...
func (c *Authentications) GetByEmail(db *gorm.DB, email string) error {
	return db.Find(c, "email = ?", email).Error
}
