package application

type Input struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"username" binding:"required"`
}