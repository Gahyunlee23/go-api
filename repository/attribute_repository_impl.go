package repository

import (
	"main-admin-api/models"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttributeRepositoryImpl struct {
	db *gorm.DB
}

func NewAttributeRepositoryImpl(db *gorm.DB) *AttributeRepositoryImpl {
	return &AttributeRepositoryImpl{db: db}
}

func (r *AttributeRepositoryImpl) Create(Attribute *models.Attribute) error {
	return r.db.Create(Attribute).Error
}

func (r *AttributeRepositoryImpl) GetByID(id uint) (*models.Attribute, error) {
	var attribute models.Attribute
	if err := r.db.First(&attribute, id).Error; err != nil {
		return nil, err
	}
	return &attribute, nil
}

func (r *AttributeRepositoryImpl) GetAll(ctx *gin.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "code", "name")).Find(&attributes).Error; err != nil {
		return nil, err
	}
	return attributes, nil
}

func (r *AttributeRepositoryImpl) Update(Attribute *models.Attribute) error {
	return r.db.Save(Attribute).Error
}

func (r *AttributeRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Attribute{}, id).Error
}
