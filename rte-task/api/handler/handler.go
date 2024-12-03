package handler

import "github.com/Vigneshwartt/golang-rte-task/service"

type UserHandler struct {
	DB service.UserService
}

func NewHandlerRepository(db service.UserService) UserHandler {
	return UserHandler{DB: db}
}
