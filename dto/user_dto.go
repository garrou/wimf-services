package dto

import "strings"

type UserDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// UserLoginDto represents a user during login
type UserLoginDto struct {
	Username string `json:"username" binding:"required" validate:"min:3,max:50"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
}

// UserCreateDto represents a user during register
type UserCreateDto struct {
	Username string `json:"username" binding:"required" validate:"min:3,max:50"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:":min:8,max:50"`
}

// TrimSpace spaces on user information
func (u *UserCreateDto) TrimSpace() {
	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)
	u.Confirm = strings.TrimSpace(u.Confirm)
}

// IsValid checks if user information is valid
func (u *UserCreateDto) IsValid() bool {
	return len(u.Password) >= 8 && u.Password == u.Confirm && len(u.Username) >= 3
}

type UsernameDto struct {
	UserId   string
	Username string `json:"username" binding:"required" validate:"min:3,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:"min:3,max:50"`
}

// TrimSpace spaces on user information
func (u *UsernameDto) TrimSpace() {
	u.Username = strings.TrimSpace(u.Username)
	u.Confirm = strings.TrimSpace(u.Confirm)
}

func (u *UsernameDto) IsValid() bool {
	return len(u.Username) >= 3 && u.Username == u.Confirm
}

type PasswordDto struct {
	UserId   string
	Current  string `json:"current" binding:"required"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:"min:8,max:50"`
}

// TrimSpace spaces on user information
func (p *PasswordDto) TrimSpace() {
	p.Password = strings.TrimSpace(p.Password)
	p.Confirm = strings.TrimSpace(p.Confirm)
}

// IsValid checks if user information is valid
func (p *PasswordDto) IsValid() bool {
	return len(p.Password) >= 8 && p.Password == p.Confirm
}
