package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
)

type AttributeServiceImpl struct {
	attributeRepository repository.AttributeRepositoryInterface
}

func NewAttributeServiceImpl(repository repository.AttributeRepositoryInterface) AttributeServiceInterface {
	return &AttributeServiceImpl{attributeRepository: repository}
}

func (s *AttributeServiceImpl) CreateAttribute(attribute *models.Attribute, ctx *gin.Context) error {
	var err error
	attribute.Settings, err = utils.MarshalAndAssignJSON(attribute.Settings, "settings", ctx)
	if err != nil {
		return err
	}

	if err := s.attributeRepository.Create(attribute); err != nil {
		return err
	}

	return nil
}

func (s *AttributeServiceImpl) GetAttributeByID(id uint) (*models.Attribute, error) {
	return s.attributeRepository.GetByID(id)
}

func (s *AttributeServiceImpl) GetAllAttributes(ctx *gin.Context) ([]models.Attribute, error) {
	return s.attributeRepository.GetAll(ctx)
}

func (s *AttributeServiceImpl) UpdateAttribute(attribute *models.Attribute, ctx *gin.Context) error {
	var err error
	attribute.Settings, err = utils.MarshalAndAssignJSON(attribute.Settings, "settings", ctx)
	if err != nil {
		return err
	}

	if err := s.attributeRepository.Update(attribute); err != nil {
		return err
	}

	return nil
}

func (s *AttributeServiceImpl) DeleteAttribute(id uint) error {
	return s.attributeRepository.Delete(id)
}

func (s *AttributeServiceImpl) ArchiveAttribute(id uint) error {
	return s.attributeRepository.Archive(id)
}
