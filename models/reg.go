package models

type Input struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"usernameb" binding:"required"`
}