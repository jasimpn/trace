// models/user.go
package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}