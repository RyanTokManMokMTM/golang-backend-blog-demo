//Package dao - database access object - manipulate Database
package dao

import (
	"github.com/RyanTokManMokMTM/blog-service/internal/model"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagLists(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{Model: &model.Model{CreatedBy: createBy}, Name: name, State: state}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{Model: &model.Model{ID: id, ModifiedBy: modifiedBy}, Name: name, State: state}
	//return tag.Update(d.engine)

	value := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		value["name"] = name
	}

	return tag.Update(d.engine, value)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
