package repository

import (
	"github.com/aldiaprilianto/takana/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetOrganizationHierarchy(orgStatus string) ([]entity.Organization, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetOrganizationHierarchy(orgStatus string) (data []entity.Organization, err error) {
	db := u.db

	query := `
	SELECT *
	FROM  organization o 
	WHERE o.org_status = ?
	`
	result := db.Debug().Raw(query, orgStatus).Scan(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
