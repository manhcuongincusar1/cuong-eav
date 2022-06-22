package domain

import (
	"github.com/jinzhu/gorm"
)

type EavEntityType struct {
	EntityTypeID   uint   `gorm:"column:entity_type_id;primary_key;auto_increment"`
	EntityTypeCode string `gorm:"column:entity_type_code"`
	EntityTable    string `gorm:"column:entity_table"`
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a EavEntityType) FindOne(condition interface{}) (*EavEntityType, error) {
	var model EavEntityType
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
