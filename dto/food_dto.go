package dto

import (
	"time"
)

type FoodCreateDto struct {
	Name       string    `json:"name" binding:"required"`
	CategoryId int       `json:"categoryId" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
	FreezeAt   time.Time `json:"freezeAt" binding:"required"`
	UserId     string
}

type FoodDto struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"categoryId"`
	Quantity   int       `json:"quantity"`
	AddedAt    time.Time `json:"addedAt"`
	FreezeAt   time.Time `json:"freezeAt"`
}
