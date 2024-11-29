package main

import (
	"log"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/drivers"
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/repository"
	"github.com/Vigneshwartt/golang-rte-task/routers"
	"github.com/Vigneshwartt/golang-rte-task/service"
	"github.com/gin-gonic/gin"
)

func main() {
	dbconnection := drivers.ConnectingDatabase()
	userrepo := repository.NewUserRepsoitory(dbconnection)
	userservice := service.NewUserService(userrepo)
	userHandler := handler.NewHandlerRepository(userservice)

	newrouter := gin.Default()
	routers.AuthRoutes(newrouter, userHandler)
	routers.UserRoutes(newrouter, userHandler)
	routers.AdminRouter(newrouter, userHandler)

	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", newrouter)
}
