package main

import (
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"              // assuming you're using Gin
    "github.com/gin-contrib/cors"
    swaggerFiles "github.com/swaggo/files"               // swagger static files
    ginSwagger "github.com/swaggo/gin-swagger"
    "Web-app/backend/database"
    "Web-app/backend/routes"
    docs "Web-app/backend/docs" // generated docs
)


// @title           Web App API
// @version         1.0
// @description     This is a sample server for a web app.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
// @BasePath  /

func main() {
    // Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

    // Initialize Swagger docs
	docs.SwaggerInfo.Title = "Web App API"
	docs.SwaggerInfo.Description = "A comprehensive API for a web app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8088"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

    // Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()




    // Initialize Gin router
    r := gin.Default()

    // Configure CORS
	allowedOrigins := []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3003", "http://localhost:8088"}
    
    
    // Add origins from environment variable if set
	if envOrigins := os.Getenv("ALLOWED_ORIGINS"); envOrigins != "" {
		for _, origin := range strings.Split(envOrigins, ",") {
			if trimmed := strings.TrimSpace(origin); trimmed != "" {
				allowedOrigins = append(allowedOrigins, trimmed)
			}
		}
		// log.Printf("Added CORS origins from environment: %v", envOrigins)
	}


    // Add CORS middleware
    r.Use(cors.New(cors.Config{
    AllowOrigins:     allowedOrigins,
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control"},
    ExposeHeaders:    []string{"Content-Length", "Authorization"},
    AllowCredentials: true,
    AllowWildcard:    false,          // optional – can be removed (default=false)
    MaxAge:           12 * time.Hour, // optional – default is the same
}))

    // Initialize routes
	routes.SetupRoutes(r, db)

    // Serve static files (uploaded images)
	// r.Static("/uploads", "./uploads")

    // Add Swagger documentation
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // // Example bcrypt usage (just to make the import "live")
    // hash, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), 10)
    // fmt.Println("Example hash:", string(hash))


    // Get port from environment or use default
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8088"
	}

	log.Printf("🚀 Web App API starting on port %s", port)
	log.Printf("📚 Swagger documentation available at: http://localhost:%s/swagger/index.html", port)
	log.Printf("🏥 Health check available at: http://localhost:%s/health", port)


    // Health check endpoint
    r.GET("/health", healthCheck)

   if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// HealthCheck godoc
// @Summary      Health Check
// @Description  Check if the server is up and running
// @Tags         health
// @Produce      plain
// @Success      200  {string}  string  "OK"
// @Router       /health [get]
func healthCheck(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}

