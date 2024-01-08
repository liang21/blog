package service

import (
	"context"
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
	"github.com/liang21/blog/internal/blog/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	v1.UnsafeCategoryServiceServer
	c *biz.CategoryCase
}

func NewCategoryService(c *biz.CategoryCase) *CategoryService {
	return &CategoryService{
		c: c,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *v1.CreateCategoryRequest) (*v1.Category, error) {
	categoryDO := &biz.Category{
		Name: in.GetName(),
		Desc: in.GetDesc(),
	}
	err := c.c.Create(ctx, categoryDO)
	if err != nil {
		return nil, err
	}
	return types.CategoryToStruct(categoryDO), nil
}
func (c *CategoryService) GetCategory(ctx context.Context, in *v1.GetCategoryRequest) (*v1.Category, error) {
	categoryDO, err := c.c.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return types.CategoryToStruct(categoryDO), nil
}
func (c *CategoryService) UpdateCategory(ctx context.Context, in *v1.UpdateCategoryRequest) (*v1.Category, error) {
	categoryDO := &biz.Category{
		Name: in.GetName(),
		Desc: in.GetDesc(),
	}
	err := c.c.Update(ctx, in.GetId(), categoryDO)
	if err != nil {
		return nil, err
	}
	return types.CategoryToStruct(categoryDO), nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, in *v1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := c.c.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (c *CategoryService) ListCategory(ctx context.Context, in *v1.ListCategoryRequest) (*v1.ListCategoryResponse, error) {
	categoryDOs, err := c.c.List(ctx)
	if err != nil {
		return nil, err
	}
	categorys := make([]*v1.Category, 0)
	for _, categoryDO := range categoryDOs {
		categorys = append(categorys, types.CategoryToStruct(categoryDO))
	}
	return &v1.ListCategoryResponse{
		Categories: categorys,
		Total:      0,
	}, nil
}
