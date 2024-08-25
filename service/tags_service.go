package service

import (
	"go-gin-gorm-crud-postgres/data/request"
	"go-gin-gorm-crud-postgres/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
