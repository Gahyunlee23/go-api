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
	Attribute := &models.Attribute{ID: id}
	if err := r.db.Model(Attribute).First(Attribute).Error; err != nil {
		return nil, err
	}
	return Attribute, nil
}

func (r *AttributeRepositoryImpl) GetAll(ctx *gin.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	if err := r.db.Model(&models.Attribute{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "code", "name")).Find(&attributes).Error; err != nil {
		return nil, err
	}
	return attributes, nil
}

func (r *AttributeRepositoryImpl) Update(Attribute *models.Attribute) error {
	return r.db.Model(Attribute).Updates(Attribute).Error
}

func (r *AttributeRepositoryImpl) Delete(id uint) error {
	attribute := &models.Attribute{ID: id}
	return r.db.Model(attribute).Delete(id).Error
}

func (r *AttributeRepositoryImpl) Archive(id uint) error {
	attribute := &models.Attribute{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx.Model(attribute), attribute, id)
	})
}
