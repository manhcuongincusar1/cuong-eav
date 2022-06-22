package domain

type EavEntityStore struct {
	EntityStoreID uint `gorm:"column:entity_store_id;primary_key;auto_increment"`
	EntityTypeID  int  `gorm:"column:entity_type_id"`
	StoreID       int  `gorm:"column:store_id"`
}
