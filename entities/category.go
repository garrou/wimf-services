package entities

type Category struct {
	ID   int    `gorm:"autoIncrement;"`
	Name string `gorm:"type:varchar(50);not null;"`
}
