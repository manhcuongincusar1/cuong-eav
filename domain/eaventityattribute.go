package domain

import (
	e "cuong-eav/constants/entity"
	"fmt"

	"github.com/jinzhu/gorm"
)

type EavEntityAttribute struct {
	EntityAttributeID uint `gorm:"column:eav_entity_attribute;primary_key;auto_increment"`
	AttributeGroupID  int  `gorm:"column:attribute_group_id"`
	AttributeID       int  `gorm:"column:attribute_id"`
	AttributeSetID    int  `gorm:"column:attribute_set_id"`
	EntityTypeID      int  `gorm:"column:entity_type_id"`
}

func (EavEntityAttribute) TableName() string {
	return e.EavEntityAttributeTable
}

func (a *EavEntityAttribute) Create() error {
	return db.Create(&a).Error
}

type AttributeListResponse struct {
	AttributeID    int    `gorm:"column:attribute_id"`
	BackendGateway string `gorm:"column:backend_gateway"`
	BackendField   string `gorm:"column:backend_field"`
	BackendType    string `gorm:"column:backend_type"`
	BackendTable   string `gorm:"column:backend_table"`
	FrontendInput  string `gorm:"column:frontend_input"`
	FrontendType   string `gorm:"column:frontend_type"`
	FrontendLabel  string `gorm:"column:frontend_label"`
	EntityTypeId   int    `gorm:"column:entity_type_id"`
	AttributeSetId int    `gorm:"column:attribute_set_id"`
	ResponseField  string `gorm:"column:response_field"`
	IsRequired     bool   `gorm:"column:is_required"`
	Note           string `gorm:"column:note"`
}

func GetAttributeList(attributeSet, entityType string) ([]*AttributeListResponse, error) {
	var model []*AttributeListResponse
	selectSqlInit := `%[1]s.attribute_id, %[1]s.backend_gateway, %[1]s.backend_field, %[1]s.backend_type, %[1]s.backend_table, %[1]s.frontend_input, %[1]s.response_field, %[1]s.is_required, %[1]s.frontend_type,
ent.entity_type_id, attrs.attribute_set_id`
	selectSQL := fmt.Sprintf(selectSqlInit, e.EavAttributeTable)
	if err := db.Select(selectSQL).Debug().
		Table(e.EavEntityAttributeTable).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavEntityAttributeTable, e.EavAttributeTable)).
		Joins(fmt.Sprintf("left join %s ent on %s.entity_type_id = ent.entity_type_id", e.EavEntityTypeTable, e.EavEntityAttributeTable)).
		Joins(fmt.Sprintf("left join %s attrs on %s.attribute_set_id = attrs.attribute_set_id", e.EavAttributeSetTable, e.EavEntityAttributeTable)).
		Where("ent.entity_type_code = ?", entityType).
		Where("attrs.attribute_set_name = ?", attributeSet).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}

func GetAttribute(entityType, attributeSet, attributeCode string) (*AttributeListResponse, error) {
	var model AttributeListResponse
	selectSqlInit := `%[1]s.attribute_id, %[1]s.note, %[1]s.backend_table, %[1]s.backend_type, %[1]s.backend_field`
	selectSQL := fmt.Sprintf(selectSqlInit, e.EavAttributeTable)
	if err := db.Select(selectSQL).
		Table(e.EavEntityAttributeTable).
		Joins(fmt.Sprintf("inner join %s on %s.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavEntityAttributeTable, e.EavAttributeTable)).
		Joins(fmt.Sprintf("left join %s on %s.entity_type_id = %s.entity_type_id", e.EavEntityTypeTable, e.EavEntityAttributeTable, e.EavEntityTypeTable)).
		Joins(fmt.Sprintf("left join %s on %s.attribute_set_id = %s.attribute_set_id", e.EavAttributeSetTable, e.EavEntityAttributeTable, e.EavAttributeSetTable)).
		Where(fmt.Sprintf("%s.entity_type_code = ?", e.EavEntityTypeTable), entityType).
		Where(fmt.Sprintf("%s.attribute_set_name = ?", e.EavAttributeSetTable), attributeSet).
		Where(fmt.Sprintf("%s.attribute_code = ?", e.EavAttributeTable), attributeCode).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &model, nil
}
