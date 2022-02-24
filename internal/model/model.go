// Package model model shared database field
package model

import (
	"fmt"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	//"gorm.io/driver/mysql"
)

type (
	Model struct {
		ID         uint32 `json:"id" gorm:"primary_key"`
		CreatedBy  string `json:"created_by"`
		ModifiedBy string `json:"modified_by"`
		CreatedOn  uint32 `json:"created_on"`
		ModifiedOn uint32 `json:"modified_on"`
		DeletedOn  uint32 `json:"deleted_on"`
		IsDel      uint8  `json:"is_del"`
	}
)

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.User,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, config)

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	//apply the callback function
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack) //when creating call this function
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack) //when updating call this function
	db.Callback().Delete().Replace("gorm:delete", deleteCallBack)                              // when deleting calling this function for instead

	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	//db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	//if err != nil {
	//	return nil, err
	//}
	//
	//sqlDB, err := db.DB()
	//if err != nil {
	//	return nil, err
	//}
	//
	////allow by default
	////if global.ServerSetting.RunMode == "debug"{
	////
	////}
	////
	//
	//sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

//Global Field call back function

func updateTimeStampForCreateCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		//for generate creating time stamp
		nowTime := time.Now().Unix()

		//update create time
		if createField, ok := scope.FieldByName("CreatedOn"); ok { //get colum/field by fieldName
			//find the field
			if createField.IsBlank {
				//CreateOn field is empty -> set the value
				_ = createField.Set(nowTime)
			}
		}

		//Update modify time
		if modifyField, ok := scope.FieldByName("ModifiedOn"); ok { //get colum/field by fieldName
			if modifyField.IsBlank {
				_ = modifyField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallBack(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok { //get the column by flag of gorm
		//the flag is not set to any column
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix()) //set modified time
	}

}

func deleteCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOpt string
		//get the field
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOpt = fmt.Sprintln(str) //set
			//log.Println(extraOpt)
		}

		//Find this 2 fields on scope
		deleteField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDel, hasIsDeletedOnField := scope.FieldByName("isDel")
		//Soft delete

		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDeletedOnField {
			//just update the the value
			now := time.Now().Unix()
			sql := fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDel.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //if current scope had any condition sql ,combine to current sql(eg:where ,join etc....)
				addExtraSpaceIfExist(extraOpt),                     //combine the option
			)
			scope.Raw(sql).Exec()
		} else {
			//hard delete
			sql := fmt.Sprintf("DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //combine the sql
				addExtraSpaceIfExist(extraOpt),
			)
			scope.Raw(sql).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
