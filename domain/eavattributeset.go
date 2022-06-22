package domain

import (
	"github.com/jinzhu/gorm"
)

type EavAttributeSet struct {
	AttributeSetID   uint   `gorm:"column:attribute_set_id;primary_key;auto_increment"`
	AttributeSetName string `gorm:"column:attribute_set_name"`
	EntityTypeID     int    `gorm:"column:entity_type_id"`
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a EavAttributeSet) FindOne(condition interface{}) (*EavAttributeSet, error) {
	var model EavAttributeSet
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}

// Take Usage: Get a row by Primary Key ([id] was the default)
func (a *EavAttributeSet) Take() error {
	if err := db.Take(&a).Error; err != nil {
		return err
	}
	return nil
}
