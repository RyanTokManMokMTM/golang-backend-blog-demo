//Package model - tagObject - Database service -and manipulate by dao obj
package model

import "github.com/jinzhu/gorm"

type (
	Tag struct {
		*Model
		Name  string `json:"name"`
		State uint8  `json:"state"`
	}
)

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name) //searching by the name
	}
	db = db.Where("state = ?", t.State)                                              //searching by the state
	if err := db.Model(&t).Where("id_del = ? ", 0).Count(&count).Error; err != nil { //count the record
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)
	err = db.Where("is_del = ?", 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, value interface{}) error {
	err := db.Model(t).Where("id = ? AND is_del = ?", t.Model.ID, 0).Updates(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

//GetInfo of the tag
func (t Tag) GetInfo(db *gorm.DB) (Tag, error) {
	var result Tag
	if err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&result).Error; err != nil {
		return result, err //no result
	}
	return result, nil

}
