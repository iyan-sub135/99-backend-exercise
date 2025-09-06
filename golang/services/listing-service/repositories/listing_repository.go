package repositories

import (
	"listing-service/models"

	"gorm.io/gorm"
)

type ListingRepository struct {
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) *ListingRepository {
	return &ListingRepository{db: db}
}

func (r *ListingRepository) GetListings(params models.GetListingsParams) ([]models.Listing, error) {
	var listings []models.Listing

	query := r.db.Model(&models.Listing{})

	if params.UserID != nil {
		query = query.Where("user_id = ?", *params.UserID)
	}

	offset := (params.PageNum - 1) * params.PageSize
	err := query.Order("created_at DESC").
		Limit(params.PageSize).
		Offset(offset).
		Find(&listings).Error

	return listings, err
}

func (r *ListingRepository) CreateListing(req models.CreateListingRequest) (*models.Listing, error) {
	listing := models.Listing{
		UserID:      req.UserID,
		ListingType: req.ListingType,
		Price:       req.Price,
	}

	err := r.db.Create(&listing).Error
	if err != nil {
		return nil, err
	}

	return &listing, nil
}
