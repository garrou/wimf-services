package entities

type User struct {
	ID       string `gorm:"unique;type:varchar(50);not null;"`
	Username string `gorm:"type:varchar(50);not null;"`
	Password string `gorm:"not null;type:varchar(255);not null"`
	Foods    []Food
}
