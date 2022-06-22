package domain

import (
	e "cuong-eav/constants/entity"
	"fmt"

	"github.com/jinzhu/gorm"
)

type EavAttributeOptionValue struct {
	ValueID          uint   `gorm:"column:value_id;primary_key;auto_increment"`
	OptionID         int    `gorm:"column:option_id"`
	StoreID          int    `gorm:"column:store_id"`
	Key              string `gorm:"column:key"`
	Value            string `gorm:"column:value"`
	SortOrder        int    `gorm:"column:sort_order"`
	StuffInformation string `gorm:"column:stuff_information"`
	Status           string `gorm:"column:status"`
}

func (a *EavAttributeOptionValue) Create() error {
	return db.Create(&a).Error
}

// Modify a single record based on Primary Key ([id] was the default)
func (a *EavAttributeOptionValue) Updates(data interface{}) error {
	return db.Model(a).Updates(data).Error
}

func (a EavAttributeOptionValue) DeleteBy(condition EavAttributeOptionValue) error {
	return db.Where(condition).Delete(EavAttributeOptionValue{}).Error
}

type AttributeOption struct {
	ValueID          int         `gorm:"column:value_id"`
	Value            string      `gorm:"column:value"`
	Status           string      `gorm:"column:status"`
	Sequence         int         `gorm:"column:sort_order"`
	StuffInformation string      `gorm:"column:stuff_information" json:"-"`
	ExtraAttribute   interface{} `json:"extraAttribute"`
}

func GetAttributeOptions(entityTypeCode, attributeSetName, attributeCode, key, value string, storeID int, status, order string) ([]*AttributeOption, error) {
	keyCondition := func() string {
		if key != "" {
			return e.EavAttributeOptionValueTable + ".key = '" + key + "'"
		}
		return "true"
	}()
	valueCondition := func() string {
		if value != "" {
			return e.EavAttributeOptionValueTable + ".value like '%" + value + "%'"
		}
		return "true"
	}()
	statusCondition := func() string {
		if status != "" {
			return fmt.Sprintf("%s.status = '%s'", e.EavAttributeOptionValueTable, status)
		}
		return "true"
	}()
	orderBy := func() string {
		if order != "" {
			return fmt.Sprintf("%s.sort_order %s", e.EavAttributeOptionValueTable, order)
		}
		return ""
	}()
	selectSQLInit := "%[1]s.key as value_id, %[1]s.value, %[1]s.status, %[1]s.sort_order as sort_order, %[1]s.stuff_information as stuff_information"
	selectSQL := fmt.Sprintf(selectSQLInit, e.EavAttributeOptionValueTable)
	var model []*AttributeOption
	if err := db.Select(selectSQL).
		Table(e.EavAttributeOptionValueTable).
		Joins(fmt.Sprintf("left join %s on %s.option_id = %s.option_id", e.EavAttributeOptionTable, e.EavAttributeOptionTable, e.EavAttributeOptionValueTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavEntityAttributeTable, e.EavEntityAttributeTable, e.EavAttributeOptionTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavEntityAttributeTable, e.EavAttributeTable)).
		Joins(fmt.Sprintf("left join %s on %s.entity_type_id = %s.entity_type_id", e.EavEntityTypeTable, e.EavEntityAttributeTable, e.EavEntityTypeTable)).
		Joins(fmt.Sprintf("left join %s on %s.attribute_set_id = %s.attribute_set_id", e.EavAttributeSetTable, e.EavEntityAttributeTable, e.EavAttributeSetTable)).
		Where(fmt.Sprintf("%s.entity_type_code = ?", e.EavEntityTypeTable), entityTypeCode).
		Where(fmt.Sprintf("%s.attribute_set_name = ?", e.EavAttributeSetTable), attributeSetName).
		Where(fmt.Sprintf("%s.attribute_code = ?", e.EavAttributeTable), attributeCode).
		Where(fmt.Sprintf("%s.store_id = ?", e.EavAttributeOptionValueTable), storeID).
		Where(statusCondition).
		Where(keyCondition).
		Where(valueCondition).
		Order(orderBy).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}

func GetAttributeOption(entityTypeCode, attributeSetName, attributeCode string, storeID int, status, valueID string) (*AttributeOption, error) {
	statusCondition := func() string {
		if status != "" {
			return fmt.Sprintf("%s.status = '%s'", e.EavAttributeOptionValueTable, status)
		}
		return "true"
	}()
	selectSqlInit := "%[1]s.value_id as value_id, %[1]s.value"
	selectSQL := fmt.Sprintf(selectSqlInit, e.EavAttributeOptionValueTable)
	var model AttributeOption
	if err := db.Select(selectSQL).
		Table(e.EavAttributeOptionValueTable).
		Joins(fmt.Sprintf("left join %s on %s.option_id = %s.option_id", e.EavAttributeOptionTable, e.EavAttributeOptionTable, e.EavAttributeOptionValueTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavEntityAttributeTable, e.EavEntityAttributeTable, e.EavAttributeOptionTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavEntityAttributeTable, e.EavAttributeTable)).
		Joins(fmt.Sprintf("left join %s on %s.entity_type_id = %s.entity_type_id", e.EavEntityTypeTable, e.EavEntityAttributeTable, e.EavEntityTypeTable)).
		Joins(fmt.Sprintf("left join %s on %s.attribute_set_id = %s.attribute_set_id", e.EavAttributeSetTable, e.EavEntityAttributeTable, e.EavAttributeSetTable)).
		Where(fmt.Sprintf("%s.entity_type_code = ?", e.EavEntityTypeTable), entityTypeCode).
		Where(fmt.Sprintf("%s.attribute_set_name = ?", e.EavAttributeSetTable), attributeSetName).
		Where(fmt.Sprintf("%s.attribute_code = ?", e.EavAttributeTable), attributeCode).
		Where(fmt.Sprintf("%s.store_id = ?", e.EavAttributeOptionValueTable), storeID).
		Where(fmt.Sprintf("%s.key = ?", e.EavAttributeOptionValueTable), valueID).
		Where(statusCondition).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}

func GetLastAttributeOptionItem(entityTypeCode, attributeSetName, attributeCode string, storeID int) (*AttributeOption, error) {
	var model AttributeOption
	selectSqlInit := "%[1]s.key as value_id, %[1]s.value"
	selectSQL := fmt.Sprintf(selectSqlInit, e.EavAttributeOptionValueTable)
	if err := db.Select(selectSQL).
		Table(e.EavAttributeOptionValueTable).
		Joins(fmt.Sprintf("left join %s on %s.option_id = %s.option_id", e.EavAttributeOptionTable, e.EavAttributeOptionTable, e.EavAttributeOptionValueTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavEntityAttributeTable, e.EavEntityAttributeTable, e.EavAttributeOptionTable)).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavEntityAttributeTable, e.EavAttributeTable)).
		Joins(fmt.Sprintf("left join %s on %s.entity_type_id = %s.entity_type_id", e.EavEntityTypeTable, e.EavEntityAttributeTable, e.EavEntityTypeTable)).
		Joins(fmt.Sprintf("left join %s on %s.attribute_set_id = %s.attribute_set_id", e.EavAttributeSetTable, e.EavEntityAttributeTable, e.EavAttributeSetTable)).
		Where(fmt.Sprintf("%s.entity_type_code = ?", e.EavEntityTypeTable), entityTypeCode).
		Where(fmt.Sprintf("%s.attribute_set_name = ?", e.EavAttributeSetTable), attributeSetName).
		Where(fmt.Sprintf("%s.attribute_code = ?", e.EavAttributeTable), attributeCode).
		Where(fmt.Sprintf("%s.store_id = ?", e.EavAttributeOptionValueTable), storeID).
		Order(fmt.Sprintf("cast(%s.key as unsigned) desc", e.EavAttributeOptionValueTable)).
		Limit(1).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
