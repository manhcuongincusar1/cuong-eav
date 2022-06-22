package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserEntityDecimal struct {
	ValueID     uint      `gorm:"column:value_id;primary_key;auto_increment"`
	AttributeID int       `gorm:"column:attribute_id"`
	EntityID    int       `gorm:"column:entity_id"`
	StoreID     int       `gorm:"column:store_id"`
	Value       float64   `gorm:"column:value"`
	Reference   int       `gorm:"column:reference"` // refer to the previous version of data
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (a *UserEntityDecimal) Create() error {
	return db.Create(&a).Error
}

func (a UserEntityDecimal) UpdatesBy(condition interface{}, data interface{}) error {
	return db.Model(a).Where(condition).Updates(data).Error
}

func (a UserEntityDecimal) Find(condition interface{}) ([]*UserEntityDecimal, error) {
	var model []*UserEntityDecimal
	if err := db.Where(condition).Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a UserEntityDecimal) FindOne(condition interface{}) (*UserEntityDecimal, error) {
	var model UserEntityDecimal
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
