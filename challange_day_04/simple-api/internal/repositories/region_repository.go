package repositories

import (
	"context"

	"github.com/codeid/hr-api/internal/models"
	"gorm.io/gorm"
)

type RegionRepository interface {
	FindAll(ctx context.Context) ([]models.Region, error)
	FindByID(ctx context.Context, id uint) (*models.Region, error)
	Create(ctx context.Context, region *models.Region) error
	Update(ctx context.Context, region *models.Region) error
	Delete(ctx context.Context, id uint) error
	FindAllWithCountry(ctx context.Context) ([]models.Region, error)
	FindByIDWithCountry(ctx context.Context, id uint) (*models.Region, error)
}

type regionRepository struct {
	DB *gorm.DB
}

// func NewRegionRepo implements RegionRepositoryInteface

func NewRegionRepository(dB *gorm.DB) RegionRepository {
	return &regionRepository{
		DB: dB,
	}
}

// FindByIDWithCountry implements [RegionRepository].
func (r *regionRepository) FindByIDWithCountry(ctx context.Context, id uint) (*models.Region, error) {
	//1. using value slice
	var regions *models.Region
	err := r.DB.WithContext(ctx).Preload("Countries").Find(&regions, "region_id = ?", id).Error
	return regions, err
}

// FindAllWithCountry implements [RegionRepository].
func (r *regionRepository) FindAllWithCountry(ctx context.Context) ([]models.Region, error) {
	//1. using value slice
	var regions []models.Region
	err := r.DB.WithContext(ctx).Preload("Countries").Find(&regions).Error
	return regions, err
}

// Create implements [RegionRepository].
func (r *regionRepository) Create(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Create(region).Error
}

// Delete implements [RegionRepository].
func (r *regionRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Region{}, id).Error
}

// FindAll implements [RegionRepository].
func (r *regionRepository) FindAll(ctx context.Context) ([]models.Region, error) {
	var regions []models.Region
	err := r.DB.WithContext(ctx).Find(&regions).Error
	return regions, err

}

// FindByID implements [RegionRepository].
func (r *regionRepository) FindByID(ctx context.Context, id uint) (*models.Region, error) {
	var region models.Region
	err := r.DB.WithContext(ctx).First(&region, id).Error
	if err != nil {
		return nil, err
	}
	return &region, nil
}

// Update implements [RegionRepository].
func (r *regionRepository) Update(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Save(region).Error
}
