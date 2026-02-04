package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/hr-api/internal/dto"
	"github.com/codeid/hr-api/internal/response"
	"github.com/codeid/hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service services.EmployeeServiceInterface
}

func NewEmployeeHandler(service services.EmployeeServiceInterface) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.service.GetAllEmployees(c.Request.Context())
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to fetch employees: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Success", employees)
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID: "+err.Error())
		return
	}

	employee, err := h.service.GetEmployeeByID(c.Request.Context(), int32(id))
	if err != nil {
		response.SendError(c, http.StatusNotFound, "Employee not found: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Success", employee)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.CreateEmployee(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to create employee: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusCreated, "Employee created successfully", resp)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID: "+err.Error())
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.UpdateEmployee(c.Request.Context(), int32(id), &req)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to update employee: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Employee updated successfully", resp)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID: "+err.Error())
		return
	}

	if err := h.service.DeleteEmployee(c.Request.Context(), int32(id)); err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to delete employee: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Employee deleted successfully", nil)
}
