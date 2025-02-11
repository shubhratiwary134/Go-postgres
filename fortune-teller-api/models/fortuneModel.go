package models

type Fortune struct {
	ID      uint   `gorm:"primaryKey"`
	Fortune string `gorm:"not null"`
}
