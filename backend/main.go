package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"              // assuming you're using Gin
    "github.com/swaggo/files"               // swagger static files
    ginSwagger "github.com/swaggo/gin-swagger"
    "golang.org/x/crypto/bcrypt"            // for password hashing
    _ "github.com/HikmaAnwar/full-stack-web-app/backend/docs" // generated docs
)


// @title           Full Stack Web App API
// @version         1.0
// @description     This is a sample server for a full-stack web app.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
    r := gin.Default()

    // Example bcrypt usage (just to make the import "live")
    hash, _ := bcrypt.GenerateFromPassword([]byte("mypassword"), 10)
    fmt.Println("Example hash:", string(hash))

    // Example Swagger setup
    r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

    // Health check endpoint
    r.GET("/health", healthCheck)

    log.Fatal(http.ListenAndServe(":8080", r))
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

