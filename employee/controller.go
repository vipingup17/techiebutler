package employee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var request CreateEmployeeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := createNewEmployee(request.Name, request.Position, request.Salary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating new employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
	})
}

func GetEmployeeByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var employee *Employee

	if employee, err = getEmployeeByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting employee by id"})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func UpdateEmployee(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request CreateEmployeeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := updateEmployee(id, request.Name, request.Position, request.Salary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
	})
}

func DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := deleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
	})
}

func ListAllEmployees(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset := (page - 1) * limit

	employeesList, err := getAllEmployeesFromDB(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching list of employees"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employees_list": employeesList})
}
