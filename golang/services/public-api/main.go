package main

import (
	"flag"
	"log"
	"strconv"

	"public-api/controllers"
	"public-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line flags
	port := flag.Int("port", 8000, "Port to run the server on")
	debug := flag.Bool("debug", true, "Run in debug mode")
	listingServiceURL := flag.String("listing-service", "http://localhost:6000", "Listing service URL")
	userServiceURL := flag.String("user-service", "http://localhost:7001", "User service URL")
	flag.Parse()

	// Set Gin mode based on debug flag
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize service clients
	listingService := services.NewListingService(*listingServiceURL)
	userService := services.NewUserService(*userServiceURL)

	// Initialize router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Initialize controllers
	publicController := controllers.NewPublicController(listingService, userService)

	// Routes
	api := router.Group("/public-api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.String(200, "OK")
		})
		api.GET("/listings", publicController.GetListings)
		api.POST("/users", publicController.CreateUser)
		api.POST("/listings", publicController.CreateListing)
	}

	// Start server
	serverAddr := ":" + strconv.Itoa(*port)
	log.Printf("Starting public API on port %d, debug: %t", *port, *debug)
	log.Printf("Listing service: %s", *listingServiceURL)
	log.Printf("User service: %s", *userServiceURL)

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
