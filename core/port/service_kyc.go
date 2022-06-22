package port

import (
	"cuong-eav/domain"
	"cuong-eav/domain/user"
)

// for simple: ignore Update, Delete and List
type KycService interface {
	CreateUser(user *user.User, attributeSet, entityType string) error
	GetUsers() ([]*user.User, error)

	CreateAttribute(att *domain.EavAttribute) error
	GetAttributes(attributeSet, entityType string) ([]*domain.AttributeListResponse, error)
}
