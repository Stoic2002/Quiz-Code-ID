package services

import (
	"context"
	"time"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/dto"
	errs "github.com/codeid/hr-api/internal/errors"
	"github.com/codeid/hr-api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

type EmployeeServiceInterface interface {
	GetAllEmployees(ctx context.Context) ([]dto.EmployeeResponse, error)
	GetEmployeeByID(ctx context.Context, id int32) (*dto.EmployeeResponse, error)
	CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	UpdateEmployee(ctx context.Context, id int32, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, id int32) error
}

type employeeService struct {
	repo     repositories.EmployeeRepository
	validate *validator.Validate
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeServiceInterface {
	return &employeeService{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s *employeeService) GetAllEmployees(ctx context.Context) ([]dto.EmployeeResponse, error) {
	employees, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.EmployeeResponse
	for _, emp := range employees {
		responses = append(responses, s.mapToResponse(emp))
	}
	return responses, nil
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id int32) (*dto.EmployeeResponse, error) {
	employee, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := s.mapToResponse(employee)
	return &resp, nil
}

func (s *employeeService) CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return nil, errs.ErrInvalidInput // Should define specific date error ideally
	}

	employee := &models.Employee{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		HireDate:     hireDate,
		JobID:        req.JobID,
		Salary:       req.Salary,
		ManagerID:    req.ManagerID,
		DepartmentID: req.DepartmentID,
	}

	if err := s.repo.Create(ctx, employee); err != nil {
		return nil, err
	}

	resp := s.mapToResponse(employee)
	return &resp, nil
}

func (s *employeeService) UpdateEmployee(ctx context.Context, id int32, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	employee, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Partial update maps
	if req.FirstName != nil {
		employee.FirstName = req.FirstName
	}
	if req.LastName != nil {
		employee.LastName = *req.LastName
	}
	if req.Email != nil {
		employee.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		employee.PhoneNumber = req.PhoneNumber
	}
	if req.JobID != nil {
		employee.JobID = *req.JobID
	}
	if req.Salary != nil {
		employee.Salary = *req.Salary
	}
	if req.ManagerID != nil {
		employee.ManagerID = req.ManagerID
	}
	if req.DepartmentID != nil {
		employee.DepartmentID = req.DepartmentID
	}

	if req.HireDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.HireDate)
		if err == nil {
			employee.HireDate = parsedDate
		}
	}

	if err := s.repo.Update(ctx, employee); err != nil {
		return nil, err
	}
	resp := s.mapToResponse(employee)
	return &resp, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *employeeService) mapToResponse(emp *models.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		EmployeeID:   emp.EmployeeID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		PhoneNumber:  emp.PhoneNumber,
		HireDate:     emp.HireDate.Format("2006-01-02"),
		JobID:        emp.JobID,
		Salary:       emp.Salary,
		ManagerID:    emp.ManagerID,
		DepartmentID: emp.DepartmentID,
	}
}
