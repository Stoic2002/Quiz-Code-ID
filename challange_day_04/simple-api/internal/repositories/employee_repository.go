package repositories

import (
	"context"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/domain/query"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]*models.Employee, error)
	FindByID(ctx context.Context, id int32) (*models.Employee, error)
	Create(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id int32) error
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{Q: query.Use(db)}
}

type employeeRepository struct {
	Q *query.Query
}

// Create implements EmployeeRepository.
func (r *employeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	return r.Q.Employee.WithContext(ctx).Create(employee)
}

// Delete implements EmployeeRepository.
func (r *employeeRepository) Delete(ctx context.Context, id int32) error {
	_, err := r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(id)).Delete(&models.Employee{})
	return err
}

// FindAll implements EmployeeRepository.
func (r *employeeRepository) FindAll(ctx context.Context) ([]*models.Employee, error) {
	return r.Q.Employee.WithContext(ctx).Find()
}

// FindByID implements EmployeeRepository.
func (r *employeeRepository) FindByID(ctx context.Context, id int32) (*models.Employee, error) {
	return r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(id)).First()
}

// Update implements EmployeeRepository.
func (r *employeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	// Updates all fields in the struct
	return r.Q.Employee.WithContext(ctx).Save(employee)
}
