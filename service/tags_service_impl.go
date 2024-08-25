package service

import "go-gin-gorm-crud-postgres/repository"

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}
