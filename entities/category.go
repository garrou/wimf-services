package entities

type Category struct {
	ID    int    `gorm:"not null;"`
	Name  string `gorm:"type:varchar(50);not null;"`
	Image string `gorm:"type:varchar(50);not null;"`
}
