package main

import (
	"gocrudapi/controllers"
	"gocrudapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.GetEnvVariables()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()
	r.GET("/employees", controllers.GetEmployees)
	r.GET("/employee/:id", controllers.GetEmployeeById)
	r.POST("/employee", controllers.PostEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployee)
	r.Run() // listen and serve on 0.0.0.0:8080
}
