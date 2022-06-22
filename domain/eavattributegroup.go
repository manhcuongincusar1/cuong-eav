package domain

import (
	e "cuong-eav/constants/entity"
	"fmt"

	"github.com/jinzhu/gorm"
)

type EavAttributeGroup struct {
	AttributeGroupID   uint   `gorm:"column:attribute_group_id;primary_key;auto_increment"`
	AttributeGroupCode string `gorm:"column:attribute_group_code"`
	AttributeGroupName string `gorm:"column:attribute_group_name"`
	AttributeSetID     int    `gorm:"column:attribute_set_id"`
	SortOrder          int    `gorm:"column:sort_order"`
}

func GetFrontendAttribute(entityType, attributeSet, attributeGroupName string) ([]*AttributeListResponse, error) {
	selectSqlInit := "%s.attribute_id, %s.frontend_input, %s.frontend_label"
	selectSQL := fmt.Sprintf(selectSqlInit, e.EavAttributeTable, e.EavAttributeTable, e.EavAttributeTable)
	var model []*AttributeListResponse
	if err := db.Select(selectSQL).
		Table(e.EavEntityAttributeTable+" a").
		Joins(fmt.Sprintf("inner join %s on a.attribute_id = %s.attribute_id", e.EavAttributeTable, e.EavAttributeTable)).
		Joins("left join "+e.EavEntityTypeTable+" ent on a.entity_type_id = ent.entity_type_id").
		Joins("left join "+e.EavAttributeSetTable+" attrs on a.attribute_set_id = attrs.attribute_set_id").
		Joins(fmt.Sprintf("left join %s on %s.attribute_group_id = a.attribute_group_id", e.EavAttributeGroupTable, e.EavAttributeGroupTable)).
		Where("ent.entity_type_code = ?", entityType).
		Where("attrs.attribute_set_name = ?", attributeSet).
		Where(fmt.Sprintf("%s.attribute_group_code = ?", e.EavAttributeGroupTable), attributeGroupName).
		Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}
