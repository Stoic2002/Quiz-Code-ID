package dto

type CreateEmployeeRequest struct {
	FirstName    *string `json:"first_name" validate:"omitempty,max=20"`
	LastName     string  `json:"last_name" validate:"required,max=25"`
	Email        string  `json:"email" validate:"required,email,max=100"`
	PhoneNumber  *string `json:"phone_number" validate:"omitempty,max=20"`
	HireDate     string  `json:"hire_date" validate:"required"` // accepting string to parse manually or strict format
	JobID        int32   `json:"job_id" validate:"required"`
	Salary       float64 `json:"salary" validate:"required,gt=0"`
	ManagerID    *int32  `json:"manager_id" validate:"omitempty"`
	DepartmentID *int32  `json:"department_id" validate:"omitempty"`
}

type UpdateEmployeeRequest struct {
	FirstName    *string  `json:"first_name" validate:"omitempty,max=20"`
	LastName     *string  `json:"last_name" validate:"omitempty,max=25"`
	Email        *string  `json:"email" validate:"omitempty,email,max=100"`
	PhoneNumber  *string  `json:"phone_number" validate:"omitempty,max=20"`
	HireDate     *string  `json:"hire_date" validate:"omitempty"`
	JobID        *int32   `json:"job_id" validate:"omitempty"`
	Salary       *float64 `json:"salary" validate:"omitempty,gt=0"`
	ManagerID    *int32   `json:"manager_id" validate:"omitempty"`
	DepartmentID *int32   `json:"department_id" validate:"omitempty"`
}

type EmployeeResponse struct {
	EmployeeID   int32   `json:"employee_id"`
	FirstName    *string `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	HireDate     string  `json:"hire_date"` // formatted date string
	JobID        int32   `json:"job_id"`
	Salary       float64 `json:"salary"`
	ManagerID    *int32  `json:"manager_id"`
	DepartmentID *int32  `json:"department_id"`
}

type EmployeeResponseWithDetails struct {
	EmployeeResponse
	// Potentially add JobName or DepartmentName here if query supports it
}
