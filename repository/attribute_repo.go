package repository

import (
	"cuong-eav/core/port"
	"cuong-eav/domain"
)

func NewKycRepository() port.KycRepo {
	return &kycRepository{}
}

type kycRepository struct {
}

func (p *kycRepository) CreateUserEntity(user *domain.UserEntity) error {
	return user.Create()
}
func (p *kycRepository) CreateUserDataInt(data *domain.UserEntityInt) error {
	return data.Create()
}
func (p *kycRepository) CreateUserDataDecimal(data *domain.UserEntityDecimal) error {
	return data.Create()
}
func (p *kycRepository) CreateUserDataVarchar(data *domain.UserEntityVarchar) error {
	return data.Create()
}
func (p *kycRepository) CreateUserDataText(data *domain.UserEntityText) error {
	return data.Create()
}

func (p *kycRepository) GetUserEntity(condition interface{}) ([]*domain.UserEntity, error) {
	return domain.UserEntity.Find(domain.UserEntity{}, condition)
}
func (p *kycRepository) GetUserDataInt(attributeId, entityId int) ([]*domain.UserEntityInt, error) {
	return domain.UserEntityInt.Find(domain.UserEntityInt{}, domain.UserEntityInt{AttributeID: attributeId, EntityID: entityId})
}
func (p *kycRepository) GetUserDataDecimal(attributeId, entityId int) ([]*domain.UserEntityDecimal, error) {
	return domain.UserEntityDecimal.Find(domain.UserEntityDecimal{}, domain.UserEntityDecimal{AttributeID: attributeId, EntityID: entityId})
}
func (p *kycRepository) GetUserDataVarchar(attributeId, entityId int) ([]*domain.UserEntityVarchar, error) {
	return domain.UserEntityVarchar.Find(domain.UserEntityVarchar{}, domain.UserEntityVarchar{AttributeID: attributeId, EntityID: entityId})
}
func (p *kycRepository) GetUserDataText(attributeId, entityId int) ([]*domain.UserEntityText, error) {
	return domain.UserEntityText.Find(domain.UserEntityText{}, domain.UserEntityText{AttributeID: attributeId, EntityID: entityId})
}

func (p *kycRepository) GetAttributeList(attributeSet, entityType string) ([]*domain.AttributeListResponse, error) {
	return domain.GetAttributeList(attributeSet, entityType)
}
func (p *kycRepository) CreateAttribute(att *domain.EavAttribute) error {
	return att.Create()
}
