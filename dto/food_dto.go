package dto

import (
	"time"
)

type FoodCreateDto struct {
	Name       string    `json:"name" binding:"required"`
	CategoryId int       `json:"categoryId" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
	Details    string    `json:"details"`
	FreezeAt   time.Time `json:"freezeAt" binding:"required"`
	UserId     string
}

type FoodUpdateDto struct {
	Id       int       `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Quantity int       `json:"quantity" binding:"required"`
	Details  string    `json:"details"`
	FreezeAt time.Time `json:"freezeAt" binding:"required"`
	UserId   string
}

type FoodDto struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"categoryId"`
	Quantity   int       `json:"quantity"`
	Details    string    `json:"details"`
	AddedAt    time.Time `json:"addedAt"`
	FreezeAt   time.Time `json:"freezeAt"`
}
