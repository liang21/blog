package service

import (
	"context"
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
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
	return nil, nil
}

func (a *ArticleCategoryService) UpdateArticleCategory(ctx context.Context, in *v1.UpdateArticleCategoryRequest) (*v1.ArticleCategory, error) {
	return nil, nil
}

func (a *ArticleCategoryService) DeleteArticleCategory(ctx context.Context, in *v1.DeleteArticleCategoryRequest) (*emptypb.Empty, error) {
	return nil, nil
}
