package domain

import (
	"fmt"

	e "cuong-eav/constants/entity"
	s "cuong-eav/constants/standard"

	"github.com/jinzhu/gorm"
)

type EavAttribute struct {
	AttributeID    uint   `gorm:"column:attribute_id;primary_key;auto_increment"`
	AttributeCode  string `gorm:"column:attribute_code"`
	EntityTypeID   int    `gorm:"column:entity_type_id"`
	BackendGateway string `gorm:"column:backend_gateway"`
	BackendModel   string `gorm:"column:backend_model"`
	BackendField   string `gorm:"column:backend_field"`
	BackendType    string `gorm:"column:backend_type"`
	BackendTable   string `gorm:"column:backend_table"`
	FrontendInput  string `gorm:"column:frontend_input"`
	FrontendType   string `gorm:"column:frontend_type"`
	FrontendLabel  string `gorm:"column:frontend_label"`
	ResponseField  string `gorm:"column:response_field"`
	IsRequired     bool   `gorm:"column:is_required"`
	Note           string `gorm:"column:note"`
}

type EvaAttributeResponse struct {
	AttributeID    uint   `gorm:"column:attribute_id;primary_key;auto_increment" json:"attribute_id"`
	AttributeCode  string `gorm:"column:attribute_code" json:"attribute_code"`
	EntityTypeID   int    `gorm:"column:entity_type_id" json:"entityType_id"`
	BackendGateway string `gorm:"column:backend_gateway" json:"backend_gateway"`
	BackendModel   string `gorm:"column:backend_model" json:"backend_model"`
	BackendField   string `gorm:"column:backend_field" json:"backend_field"`
	BackendType    string `gorm:"column:backend_type" json:"backend_type"`
	BackendTable   string `gorm:"column:backend_table" json:"backend_table"`
	FrontendInput  string `gorm:"column:frontend_input" json:"frontend_input"`
	FrontendLabel  string `gorm:"column:frontend_label" json:"frontend_label"`
	ResponseField  string `gorm:"column:response_field" json:"response_field"`
	IsRequired     bool   `gorm:"column:is_required" json:"is_required"`
	Note           string `gorm:"column:note" json:"note"`
}

func SearchEvaAttribute(AttCode string) ([]*EvaAttributeResponse, error) {
	AttCodeCondition := func() string {
		if AttCode != "" {
			return e.EavAttributeTable + ".attribute_code like '%" + AttCode + "%'"
		}
		return "true"
	}()
	selectSQLInit := `%[1]s.attribute_id, %[1]s.attribute_code, %[1]s.entity_type_id, 
					  %[1]s.backend_gateway, %[1]s.backend_model, %[1]s.backend_field, 
					  %[1]s.backend_type ,%[1]s.backend_table, %[1]s.frontend_input,
					  %[1]s.frontend_label, %[1]s.response_field, %[1]s.is_required, %[1]s.note`
	selectSQL := fmt.Sprintf(selectSQLInit, e.EavAttributeTable)
	var data []*EvaAttributeResponse
	if err := db.Select(selectSQL).
		Table(e.EavAttributeTable).
		Where(AttCodeCondition).
		Order(fmt.Sprintf("%s.attribute_id  %s", e.EavAttributeTable, s.OrderAsc)).
		Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return data, nil
}

func (a EavAttribute) Find(condition interface{}) ([]*EavAttribute, error) {
	var model []*EavAttribute
	if err := db.Where(condition).Find(&model).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return model, nil
}
func (a *EavAttribute) Create() error {
	return db.Create(&a).Error
}

// func (a *EavAttribute) Updates(data interface{}) error {
// 	return db.Model(a).Updates(data).Error
// }
func (a EavAttribute) UpdatesBy(condition interface{}, data interface{}) error {
	return db.Model(a).Where(condition).Updates(data).Error
}

func (a *EavAttribute) Delete() error {
	return db.Delete(&a).Error
}
func (a EavAttribute) FindOne(condition interface{}) (*EavAttribute, error) {
	var model EavAttribute
	if err := db.Table(e.EavAttributeTable).
		Where(condition).
		First(&model).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &model, nil
}

type EavAttDetailsResponse struct {
	AttributeID    uint   `gorm:"column:attribute_id;primary_key;auto_increment" json:"attribute_id"`
	AttributeCode  string `gorm:"column:attribute_code" json:"attribute_code"`
	EntityTypeID   int    `gorm:"column:entity_type_id" json:"entityType_id"`
	BackendGateway string `gorm:"column:backend_gateway" json:"backend_gateway"`
	BackendModel   string `gorm:"column:backend_model" json:"backend_model"`
	BackendField   string `gorm:"column:backend_field" json:"backend_field"`
	BackendType    string `gorm:"column:backend_type" json:"backend_type"`
	BackendTable   string `gorm:"column:backend_table" json:"backend_table"`
	FrontendInput  string `gorm:"column:frontend_input" json:"frontend_input"`
	FrontendLabel  string `gorm:"column:frontend_label" json:"frontend_label"`
	ResponseField  string `gorm:"column:response_field" json:"response_field"`
	IsRequired     bool   `gorm:"column:is_required" json:"is_required"`
	Note           string `gorm:"column:note" json:"note"`
}

func DetailsEavAttribute(Id uint) (*EavAttDetailsResponse, error) {
	selectSQLInit := `%[1]s.attribute_id, %[1]s.attribute_code, %[1]s.entity_type_id, 
					  %[1]s.backend_gateway, %[1]s.backend_model, %[1]s.backend_field, 
					  %[1]s.backend_type ,%[1]s.backend_table, %[1]s.frontend_input,
					  %[1]s.frontend_label, %[1]s.response_field, %[1]s.is_required, %[1]s.note`
	selectSQL := fmt.Sprintf(selectSQLInit, e.EavAttributeTable)
	var data EavAttDetailsResponse
	if err := db.Select(selectSQL).
		Table(e.EavAttributeTable).
		Where(fmt.Sprintf("%s.attribute_id = ?", e.EavAttributeTable), Id).
		Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &data, nil
}
