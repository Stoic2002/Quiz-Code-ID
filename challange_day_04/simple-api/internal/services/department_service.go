package services

import (
	"context"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/dto"
	errs "github.com/codeid/hr-api/internal/errors"
	"github.com/codeid/hr-api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

// DepartmentServiceInterface defines the service interface
type DepartmentServiceInterface interface {
	Create(ctx context.Context, req *dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.DepartmentResponse, error)
	GetAll(ctx context.Context) ([]dto.DepartmentResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateDepartmentRequest) (*dto.DepartmentResponse, error)
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]dto.DepartmentResponse, error)
}

type departmentService struct {
	repo     repositories.DepartmentRepository
	validate *validator.Validate
}

func NewDepartmentService(repo repositories.DepartmentRepository) DepartmentServiceInterface {
	return &departmentService{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Create creates a new department
func (s *departmentService) Create(ctx context.Context, req *dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	department := &models.Department{
		DepartmentName: req.DepartmentName,
	}

	if err := s.repo.Create(ctx, department); err != nil {
		return nil, err
	}

	return &dto.DepartmentResponse{
		DepartmentID:   uint(department.DepartmentID),
		DepartmentName: department.DepartmentName,
	}, nil
}

// FindByID gets a department by ID
func (s *departmentService) FindByID(ctx context.Context, id uint) (*dto.DepartmentResponse, error) {
	dept, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.DepartmentResponse{
		DepartmentID:   uint(dept.DepartmentID),
		DepartmentName: dept.DepartmentName,
	}, nil
}

// GetAll gets all departments
func (s *departmentService) GetAll(ctx context.Context) ([]dto.DepartmentResponse, error) {
	depts, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.DepartmentResponse
	for _, dept := range depts {
		responses = append(responses, dto.DepartmentResponse{
			DepartmentID:   uint(dept.DepartmentID),
			DepartmentName: dept.DepartmentName,
		})
	}
	return responses, nil
}

// Update updates a department
func (s *departmentService) Update(ctx context.Context, id uint, req *dto.UpdateDepartmentRequest) (*dto.DepartmentResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	dept, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map DTO to model (partial update)
	if req.DepartmentName != nil {
		dept.DepartmentName = *req.DepartmentName
	}

	if err := s.repo.Update(ctx, dept); err != nil {
		return nil, err
	}

	return &dto.DepartmentResponse{
		DepartmentID:   uint(dept.DepartmentID),
		DepartmentName: dept.DepartmentName,
	}, nil
}

// Delete deletes a department by ID
func (s *departmentService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// SearchByName searches departments by name
func (s *departmentService) SearchByName(ctx context.Context, name string) ([]dto.DepartmentResponse, error) {
	if name == "" {
		return s.GetAll(ctx)
	}

	depts, err := s.repo.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var responses []dto.DepartmentResponse
	for _, dept := range depts {
		responses = append(responses, dto.DepartmentResponse{
			DepartmentID:   uint(dept.DepartmentID),
			DepartmentName: dept.DepartmentName,
		})
	}
	return responses, nil
}
