package main

import (
	"gocrudapi/controllers"
	"gocrudapi/initializers"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// For CORS: go get github.com/gin-contrib/cors
func init() {
	initializers.GetEnvVariables()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "X-HTTP-Method-Override", "Content-Type", "Authorization", "Accept", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"},

		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/employees", controllers.GetEmployees)
	r.GET("/employee/:id", controllers.GetEmployeeById)
	r.POST("/employee", controllers.PostEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployee)
	r.Run() // listen and serve on 0.0.0.0:8080

}
