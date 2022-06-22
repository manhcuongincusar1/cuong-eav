package port

import (
	"cuong-eav/domain"
)

type KycRepo interface {
	CreateUserEntity(user *domain.UserEntity) error
	CreateUserDataInt(data *domain.UserEntityInt) error
	CreateUserDataDecimal(data *domain.UserEntityDecimal) error
	CreateUserDataVarchar(data *domain.UserEntityVarchar) error
	CreateUserDataText(data *domain.UserEntityText) error

	GetUserEntity(condition interface{}) ([]*domain.UserEntity, error)
	GetUserDataInt(attributeId, entityId int) ([]*domain.UserEntityInt, error)
	GetUserDataDecimal(attributeId, entityId int) ([]*domain.UserEntityDecimal, error)
	GetUserDataVarchar(attributeId, entityId int) ([]*domain.UserEntityVarchar, error)
	GetUserDataText(attributeId, entityId int) ([]*domain.UserEntityText, error)

	GetAttributeList(attributeSet, entityType string) ([]*domain.AttributeListResponse, error)
	CreateAttribute(att *domain.EavAttribute) error
}
