package repositories

import (
	"context"

	"github.com/codeid/hr-api/internal/domain/models"
	"github.com/codeid/hr-api/internal/domain/query"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindAll(ctx context.Context) ([]*models.Department, error)
	FindByID(ctx context.Context, id uint) (*models.Department, error)
	Create(ctx context.Context, department *models.Department) error
	Update(ctx context.Context, department *models.Department) error
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]*models.Department, error)
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{Q: query.Use(db)}
}

type departmentRepository struct {
	Q *query.Query
}

// SearchByName implements DepartmentRepository.
func (d *departmentRepository) SearchByName(ctx context.Context, name string) ([]*models.Department, error) {
	departments, err := d.Q.Department.WithContext(ctx).
		Where(d.Q.Department.DepartmentName.Like("%" + name + "%")).Find()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// Create implements DepartmentRepository.
func (d *departmentRepository) Create(ctx context.Context, department *models.Department) error {
	return d.Q.Department.WithContext(ctx).Create(department)
}

// Delete implements DepartmentRepository.
func (d *departmentRepository) Delete(ctx context.Context, id uint) error {
	_, err := d.Q.Department.WithContext(ctx).Where(d.Q.Department.DepartmentID.Eq(int32(id))).Delete(&models.Department{})
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements DepartmentRepository.
func (d *departmentRepository) FindAll(ctx context.Context) ([]*models.Department, error) {
	var departments []*models.Department
	departments, err := d.Q.Department.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return departments, err
}

// FindByID implements DepartmentRepository.
func (d *departmentRepository) FindByID(ctx context.Context, id uint) (*models.Department, error) {
	department, err := d.Q.Department.WithContext(ctx).Where(d.Q.Department.DepartmentID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}
	return department, nil
}

// Update implements DepartmentRepository.
func (d *departmentRepository) Update(ctx context.Context, department *models.Department) error {
	return d.Q.Department.WithContext(ctx).Save(department)
}
