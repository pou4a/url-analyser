package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "log"
    "url-analyser-backend/database"
    "url-analyser-backend/handlers"
)

func main() {
    database.EnsureDatabase()
    db := database.InitDB()
    database.Migrate(db)

    router := gin.Default()

    // Enable CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Front-end origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }))

    // Health check route
    router.GET("/healthCheck", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "checked!"})
    })

    // URL routes
    router.POST("/urls", handlers.AddURL(db))
    router.GET("/urls", handlers.ListURLs(db))
    router.POST("/crawl", handlers.CrawlURLHandler(db))
    router.GET("/crawl-results", handlers.GetCrawlResults(db))
	
    log.Println("Server running at http://localhost:8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}