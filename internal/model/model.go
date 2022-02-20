// Package model model shared database field
package model

import (
	"fmt"
	"github.com/RyanTokManMokMTM/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Model struct {
		ID         uint32 `json:"id" gorm:"primary_key"`
		CreateBy   string `json:"create_by"`
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
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	//allow by default
	//if global.ServerSetting.RunMode == "debug"{
	//
	//}
	//

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
