package service

import (
	"context"
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/repository"
)

type CategoryServiceimpl struct {
	CategoryRespository repository.CategoryRepository
	DB                  *sql.DB
}

func (service *CategoryServiceimpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRespository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceimpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRespository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.CategoryRespository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceimpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRespository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRespository.Delete(ctx, tx, category)
}

func (service *CategoryServiceimpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRespository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceimpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRespository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
