package main

import (
	"log"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/drivers"
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/repository"
	"github.com/Vigneshwartt/golang-rte-task/routers"
	"github.com/Vigneshwartt/golang-rte-task/service"
)

func main() {
	dbconnection := drivers.ConnectingDatabase()
	userrepo := repository.NewUserRepsoitory(dbconnection)
	userservice := service.NewUserService(userrepo)
	userHandler := handler.NewHandlerRepository(userservice)
	r := routers.IntializeRouter(userHandler)
	// err := r.Run(":8080")
	// if err != nil {
	// 	fmt.Println("Error Occured", err)
	// 	return
	// }
	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", r)
}
