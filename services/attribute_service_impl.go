package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"

	"github.com/gin-gonic/gin"
)

type AttributeServiceImpl struct {
	attributeRepository repository.AttributeRepositoryInterface
}

func NewAttributeServiceImpl(repository repository.AttributeRepositoryInterface) AttributeServiceInterface {
	return &AttributeServiceImpl{attributeRepository: repository}
}

func (s *AttributeServiceImpl) CreateAttribute(attribute *models.Attribute) error {
	return s.attributeRepository.Create(attribute)
}

func (s *AttributeServiceImpl) GetAttributeByID(id uint) (*models.Attribute, error) {
	return s.attributeRepository.GetByID(id)
}

func (s *AttributeServiceImpl) GetAllAttributes(ctx *gin.Context) ([]models.Attribute, error) {
	return s.attributeRepository.GetAll(ctx)
}

func (s *AttributeServiceImpl) UpdateAttribute(attribute *models.Attribute) error {
	return s.attributeRepository.Update(attribute)
}

func (s *AttributeServiceImpl) DeleteAttribute(id uint) error {
	return s.attributeRepository.Delete(id)
}
