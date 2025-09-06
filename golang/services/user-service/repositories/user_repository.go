package repositories

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers(params models.GetUsersParams) ([]models.User, error) {
	var users []models.User

	offset := (params.PageNum - 1) * params.PageSize
	err := r.db.Model(&models.User{}).
		Order("created_at DESC").
		Limit(params.PageSize).
		Offset(offset).
		Find(&users).Error

	return users, err
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	user := models.User{
		Name: req.Name,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
