package main

import (
	"fmt"
	"log"
	"net/http"
	"weKnow/adapter"
	"weKnow/config"
	"weKnow/controller"
	"weKnow/db"
	"weKnow/middleware"
	"weKnow/repository"
	"weKnow/router"
	"weKnow/service"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	repo := repository.NewRepository(db.NewDataBase(config), adapter.NewAdapter())
	s := service.NewService(repo.(repository.KnownRepository), *config)
	ctrl := controller.NewController(s.(service.KnownService))
	// j := job.NewJob(s)
	// err = j.ScheduleJobs()
	// if err != nil {
	// 	fmt.Println("Error scheduling job:", err)
	// 	return
	// }
	// j.StartScheduler()
	r := router.SetupRouter(ctrl.(controller.KnownController))
	mr := middleware.CorsAndLoggingMiddleware(r)
	fmt.Println("Started on port", config.App.Port)
	if err := http.ListenAndServe(":"+fmt.Sprint(config.App.Port), mr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
