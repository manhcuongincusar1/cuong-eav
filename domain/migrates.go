package domain

import e "cuong-eav/constants/entity"

type Migrations struct {
	Id        uint   `gorm:"column:id;primary_key;auto_increment"`
	Migration string `gorm:"column:migration"`
	Batch     int    `gorm:"column:batch"`
}

func (a *Migrations) Migrate1Up() {

	db.Create(&EavEntityType{EntityTypeID: 1, EntityTypeCode: e.EntityTypeUser, EntityTable: e.UserEntity})
	db.Create(&EavAttributeSet{AttributeSetID: 1, AttributeSetName: e.AttributeSetDefault, EntityTypeID: 1})
	db.Create(&EavEntityAttribute{
		EntityAttributeID: 1,
		AttributeGroupID:  1,
		AttributeSetID:    1,
		AttributeID:       1,
		EntityTypeID:      1,
	})

	db.Create(&EavAttribute{AttributeID: 1, EntityTypeID: 1, AttributeCode: "marital_status", BackendModel: "", BackendField: "", BackendType: "string", BackendTable: "", FrontendInput: "marital_status", FrontendLabel: "Marital Status", IsRequired: true})
}
