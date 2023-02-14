package controllers

import (
	"gocrudapi/initializers"
	"gocrudapi/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	initializers.DB.Find(&employees)

	c.JSON(200, gin.H{
		"employees": employees,
	})
}

func PostEmployee(c *gin.Context) {

	var body struct {
		Name  string
		Email string
	}

	c.Bind(&body)

	employee := models.Employee{Name: body.Name, Email: body.Email}

	result := initializers.DB.Create(&employee)

	if result.Error != nil {
		log.Fatal("Could not be created")
	}

	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	initializers.DB.First(&employee, id)
	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee

	var body struct {
		Name  string
		Email string
	}

	c.Bind(&body)
	id := c.Param("id")
	initializers.DB.First(&employee, id)

	initializers.DB.Model(&employee).Updates(models.Employee{Name: body.Name, Email: body.Email})
	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	initializers.DB.Delete(&employee, id)
	c.JSON(200, gin.H{
		"employee": "The record has been deleted",
	})
}
