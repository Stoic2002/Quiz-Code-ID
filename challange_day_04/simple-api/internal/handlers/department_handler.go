package handlers

import (
	"net/http"
	"strconv"

	"github.com/codeid/hr-api/internal/dto"
	"github.com/codeid/hr-api/internal/response"
	"github.com/codeid/hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

// DepartmentHandler handles HTTP requests for departments
type DepartmentHandler struct {
	service services.DepartmentServiceInterface
}

// NewDepartmentHandler creates a new handler instance
func NewDepartmentHandler(svc services.DepartmentServiceInterface) *DepartmentHandler {
	return &DepartmentHandler{
		service: svc,
	}
}

// CreateDepartment handles POST /departments
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to create department: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusCreated, "Department created successfully", resp)
}

// GetDepartmentByID handles GET /departments/:id
func (h *DepartmentHandler) GetDepartmentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.FindByID(c.Request.Context(), uint(id))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Department not found: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "Department created successfully", resp)
}

// GetAllDepartments handles GET /departments
func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	resp, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch departments: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// UpdateDepartment handles PUT /departments/:id
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch departments: "+err.Error())
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch departments: "+err.Error())
		return
	}

	resp, err := h.service.Update(c.Request.Context(), uint(id), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch departments: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}

// DeleteDepartment handles DELETE /departments/:id
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.SendError(c, http.StatusNotFound, "Data not found:"+err.Error())
		return
	}
	response.SendResponse(c, http.StatusNoContent, "", nil)
}

// SearchDepartments handles GET /departments/search?q=...
func (h *DepartmentHandler) SearchDepartments(c *gin.Context) {
	name := c.Query("q")
	if name == "" {
		response.SendError(c, http.StatusBadRequest, "search query required: ")
		return
	}

	resp, err := h.service.SearchByName(c.Request.Context(), name)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch departments: "+err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "", resp)
}
