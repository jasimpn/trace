package models

type Repo struct{
	ID   uint   `gorm:"primaryKey"`
	Name string	
}