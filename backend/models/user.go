package models

import (
	"time"

	"gorm.io/gorm"
)

// ANCHOR - User model
type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"first_name" form:"first_name"`
	LastName  string         `json:"last_name" form:"last_name"`
	Age       int            `json:"age" form:"age"`
	Email     string         `json:"email" form:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	File      File           `gorm:"polymorphic:Table;polymorphicValue:users" json:"file"`
}

// func (User) Seed(db *gorm.DB) {
// 	employeeWarning := []User{
// 		{ID: 1, FirstName: "Nasuh", LastName: "Yücel", Age: 25, Email: "nasuhyc@gmail.com"},
// 	}
// 	db.Create(&employeeWarning)
// }
