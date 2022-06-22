package domain

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type UserEntity struct {
	Id          int          `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Uuid        string       `json:"uuid"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	LegalName   string       `json:"legal_name"`
	Email       string       `json:"email"`
	PhoneNumber string       `json:"phone_number"`
	Gender      int          `json:"gender"`
	Password    string       `json:"password"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

func (a *UserEntity) Create() error {
	return db.Create(&a).Error
}

func (a UserEntity) UpdatesBy(condition interface{}, data interface{}) error {
	return db.Model(a).Where(condition).Updates(data).Error
}

func (a UserEntity) Find(condition interface{}) ([]*UserEntity, error) {
	var model []*UserEntity
	if err := db.Where(condition).Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a UserEntity) FindOne(condition interface{}) (*UserEntity, error) {
	var model UserEntity
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
