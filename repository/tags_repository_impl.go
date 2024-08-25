package repository

import (
	"go-gin-gorm-crud-postgres/helper"
	"go-gin-gorm-crud-postgres/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
