package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserEntityText struct {
	ValueID     uint      `gorm:"column:value_id;primary_key;auto_increment"`
	AttributeID int       `gorm:"column:attribute_id"`
	EntityID    int       `gorm:"column:entity_id"`
	StoreID     int       `gorm:"column:store_id"`
	Value       string    `gorm:"column:value"`
	Reference   int       `gorm:"column:reference"` // refer to the previous version of data
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (a *UserEntityText) Create() error {
	return db.Create(&a).Error
}

func (a UserEntityText) UpdatesBy(condition interface{}, data interface{}) error {
	return db.Model(a).Where(condition).Updates(data).Error
}

func (a UserEntityText) Find(condition interface{}) ([]*UserEntityText, error) {
	var model []*UserEntityText
	if err := db.Where(condition).Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a UserEntityText) FindOne(condition interface{}) (*UserEntityText, error) {
	var model UserEntityText
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
