package service

import (
	"context"
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
	"github.com/liang21/blog/internal/blog/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleCategoryService struct {
	v1.UnimplementedArticleCategoryServiceServer
	ac *biz.ArticleCategoryCase
}

func NewArticleCategoryService(ac *biz.ArticleCategoryCase) *ArticleCategoryService {
	return &ArticleCategoryService{
		ac: ac,
	}
}

func (a *ArticleCategoryService) CreateArticleCategory(ctx context.Context, in *v1.CreateArticleCategoryRequest) (*v1.ArticleCategory, error) {
	articleCategoryDO := &biz.ArticleCategory{
		ArticleId:  in.GetArticleId(),
		CategoryId: in.GetCategoryId(),
	}
	err := a.ac.Create(ctx, articleCategoryDO)
	if err != nil {
		return nil, err
	}
	return types.ArticleCategoryToStruct(articleCategoryDO), nil
}

func (a *ArticleCategoryService) UpdateArticleCategory(ctx context.Context, in *v1.UpdateArticleCategoryRequest) (*v1.ArticleCategory, error) {
	articleCategory, err := a.ac.Get(ctx, in.GetArticleId())
	if err != nil {
		return nil, err
	}
	articleCategory.CategoryId = in.GetCategoryId()
	articleCategory.ArticleId = in.GetArticleId()
	err = a.ac.Update(ctx, articleCategory.Id, articleCategory)
	if err != nil {
		return nil, err
	}
	return types.ArticleCategoryToStruct(articleCategory), nil
}

func (a *ArticleCategoryService) DeleteArticleCategory(ctx context.Context, in *v1.DeleteArticleCategoryRequest) (*emptypb.Empty, error) {
	err := a.ac.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a *ArticleCategoryService) GetArticleCategory(ctx context.Context, in *v1.GetArticleCategoryRequest) (*v1.ArticleCategory, error) {
	articleCategory, err := a.ac.Get(ctx, in.GetArticleId())
	if err != nil {
		return nil, err
	}
	return types.ArticleCategoryToStruct(articleCategory), nil
}
