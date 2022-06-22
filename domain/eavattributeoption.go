package domain

import "github.com/jinzhu/gorm"

type EavAttributeOption struct {
	OptionID    uint `gorm:"column:option_id;primary_key;auto_increment"`
	AttributeID int  `gorm:"column:attribute_id"`
}

func (a *EavAttributeOption) Create() error {
	return db.Create(&a).Error
}

// Usage: You sure that the result of query only 1 row returned, this function recommended
func (a EavAttributeOption) FindOne(condition interface{}) (*EavAttributeOption, error) {
	var model EavAttributeOption
	if err := db.Where(condition).Take(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
