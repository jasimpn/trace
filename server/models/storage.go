package models

type Storage struct{
	ID   uint   `gorm:"primaryKey"`
	Name string	
}