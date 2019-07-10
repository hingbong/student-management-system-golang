package main

import (
	"net/http"
	"system-management-system/main/models"
	"system-management-system/main/routes"
	"time"
)

func main() {
	// err is already nil
	_, _ = models.InitDB()

	// Listen and Server in 0.0.0.0:8080
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	routes.SetupRouter(s)
}
