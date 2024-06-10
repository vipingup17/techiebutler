package main

import (
	"techiebutler/employee"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/employee/create", employee.CreateEmployee)
	router.GET("/employee/get/:id", employee.GetEmployeeByID)
	router.PUT("/employee/update/:id", employee.UpdateEmployee)
	router.DELETE("/employee/delete/:id", employee.DeleteEmployee)

	router.GET("/employee/list-all-employees", employee.ListAllEmployees)
}
