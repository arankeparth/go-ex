package main

import (
	"context"
	"log"
	"net/http"
	"runtime"

	_ "net/http/pprof" // Import pprof for side effects

	"go-ex/db"

	"go-ex/config"
	"go-ex/routers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	ctx := context.Background()

	config.LoadEnv()

	log.Printf("Creating Database connection.")
	db.InitializeDatabase(ctx)

	log.Printf("Setting up Routers")
	routers.SetUpUserRouters(r)

	return r
}

func main() {
	go func() {
		runtime.SetCPUProfileRate(1000000)
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe(":6060", nil)) // Default pprof routes served here
	}()
	r := setupRouter()
	// Listen and serve on 0.0.0.0:8080
	port := config.GetEnv("PORT")
	r.Run(port)
}
