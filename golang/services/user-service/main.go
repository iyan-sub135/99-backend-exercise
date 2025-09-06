package main

import (
	"flag"
	"log"
	"strconv"

	"user-service/controllers"
	"user-service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 7000, "Port to run the server on")
	debug := flag.Bool("debug", true, "Run in debug mode")
	flag.Parse()

	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	userController := controllers.NewUserController(db)

	// Routes
	v1 := router.Group("/")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.String(200, "OK")
		})
		v1.GET("/users", userController.GetUsers)
		v1.GET("/users/:id", userController.GetUserByID)
		v1.POST("/users", userController.CreateUser)
	}

	// Start server
	serverAddr := ":" + strconv.Itoa(*port)
	log.Printf("Starting user service on port %d, debug: %t", *port, *debug)

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
