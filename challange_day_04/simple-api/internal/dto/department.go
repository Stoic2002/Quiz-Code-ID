package dto

// CreateDepartmentRequest for creating a department
type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" validate:"required,min=2,max=100"` // Validation tags
}

// UpdateDepartmentRequest for partial update
type UpdateDepartmentRequest struct {
	DepartmentName *string `json:"department_name,omitempty" validate:"omitempty,min=2,max=100"`
}

// DepartmentResponse for output (hide internal fields if needed)
type DepartmentResponse struct {
	DepartmentID   uint   `json:"department_id"`
	DepartmentName string `json:"department_name"`
}
