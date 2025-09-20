package main

import (
	"fmt"
	"log"
	"net/http"
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
	repo := repository.NewRepository(db.NewDataBase(config))
	s := service.NewService(repo)
	ctrl := controller.NewController(s)
	// j := job.NewJob(s)
	// err = j.ScheduleJobs()
	// if err != nil {
	// 	fmt.Println("Error scheduling job:", err)
	// 	return
	// }
	// j.StartScheduler()
	r := router.SetupRouter(ctrl)
	mr := middleware.CorsAndLoggingMiddleware(r)
	fmt.Println("Started on port", config.App.Port)
	if err := http.ListenAndServe(":"+fmt.Sprint(config.App.Port), mr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
