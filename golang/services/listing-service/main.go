package main

import (
	"flag"
	"log"
	"strconv"

	"listing-service/controllers"
	"listing-service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 6000, "Port to run the server on")
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

	listingController := controllers.NewListingController(db)

	// Routes
	v1 := router.Group("/")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.String(200, "OK")
		})
		v1.GET("/listings", listingController.GetListings)
		v1.POST("/listings", listingController.CreateListing)
	}

	// Start server
	serverAddr := ":" + strconv.Itoa(*port)
	log.Printf("Starting listing service on port %d, debug: %t", *port, *debug)

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
