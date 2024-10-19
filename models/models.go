package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"gin_web_api/pkg/setting"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"time"
)

var db *gorm.DB

type Model struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
	IsDelete   int       `json:"is_delete"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// 连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("models.Setup Connect Poll Inital err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
//func CloseDB() {
//	defer db.Close()
//}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
//func updateTimeStampForCreateCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		nowTime := time.Now().Unix()
//		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
//			if createTimeField.IsBlank {
//				createTimeField.Set(nowTime)
//			}
//		}
//
//		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
//			if modifyTimeField.IsBlank {
//				modifyTimeField.Set(nowTime)
//			}
//		}
//	}
//}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}

// deleteCallback will set `DeletedOn` where deleting
//func deleteCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
//
//		if !scope.Search.Unscoped && hasDeletedOnField {
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v%v%v",
//				scope.QuotedTableName(),
//				scope.Quote(deletedOnField.DBName),
//				scope.AddToVars(time.Now().Unix()),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v%v%v",
//				scope.QuotedTableName(),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		}
//	}
//}

// addExtraSpaceIfExist adds a separator
//func addExtraSpaceIfExist(str string) string {
//	if str != "" {
//		return " " + str
//	}
//	return ""
//}
