package entities

import "time"

type Food struct {
	ID         int       `gorm:"autoIncrement;"`
	Name       string    `gorm:"type:varchar(255);not null;"`
	CategoryID int       `gorm:"not null;"`
	Quantity   int       `gorm:"not null;"`
	AddedAt    time.Time `gorm:"not null;"`
	FreezeAt   time.Time `gorm:"not null;"`
	Details    string    `gorm:"type:varchar(255);"`
	UserID     string    `gorm:"type:varchar(50);not null;"`
}
