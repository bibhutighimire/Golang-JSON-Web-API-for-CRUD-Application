package main

import (
	"gocrudapi/initializers"
	"gocrudapi/models"
)

func init() {
	initializers.GetEnvVariables()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{})

}
