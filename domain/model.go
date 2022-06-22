package domain

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `gorm:"column:created_on" json:"created_on"`
	ModifiedOn int `gorm:"column:modified_on" json:"modified_on"`
	DeletedOn  int `gorm:"column:deleted_on" json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"12345678",
		"127.0.0.1:3306",
		"cuong_eav"))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func Migrate() {
	db.AutoMigrate(&Migrations{})
	db.AutoMigrate(&EavAttribute{})
	db.AutoMigrate(&EavAttributeGroup{})
	db.AutoMigrate(&EavAttributeOption{})
	db.AutoMigrate(&EavAttributeSet{})
	db.AutoMigrate(&EavEntityAttribute{})
	db.AutoMigrate(&EavEntityStore{})
	db.AutoMigrate(&EavEntityType{})

	db.AutoMigrate(&UserEntity{})
	db.AutoMigrate(&UserEntityDecimal{})
	db.AutoMigrate(&UserEntityInt{})
	db.AutoMigrate(&UserEntityVarchar{})
	db.AutoMigrate(&UserEntityText{})

	_ = migrate(1)
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// Usage: Get a row by Primary Key ([id] was the default)
func Take(models interface{}) error {
	if err := db.Take(models).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func migrate(version int) error {
	var migrate Migrations
	err := db.Order("id desc").Where("id <= ?", version).Take(&migrate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	migrateNoFrom := 1
	if migrate.Id != 0 {
		if int(migrate.Id) == version {
			migrateNoFrom = version + 1
		} else {
			migrateNoFrom = int(migrate.Id) + 1
		}
	}
	for i := migrateNoFrom; i <= version; i++ {
		callableName := fmt.Sprintf("Migrate%dUp", i)
		callable := reflect.ValueOf(&Migrations{}).MethodByName(callableName)
		if !callable.IsValid() {
			log.Fatalf("\nFailed to find Registry method with name \"%s\".", callableName)
		}
		callable.Call([]reflect.Value{})

		migrate := Migrations{
			Migration: callableName,
			Batch:     1,
		}
		_ = db.Create(&migrate).Error
	}
	return nil
}

type AnonymousCode struct {
	Id   uint   `gorm:"column:id"`
	Code string `gorm:"column:code"`
}

func GetUniqueCode(code, table, column string) (*AnonymousCode, error) {
	var data AnonymousCode
	selectSQL := "id, " + column + " as code"
	err := db.Select(selectSQL).
		Table(table).
		Where(column+" = ?", code).
		Take(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &data, nil
}

type Filter struct {
	Table, Column, Value string
}
