package service

import (
	"context"
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
	"github.com/liang21/blog/internal/blog/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleService struct {
	v1.UnimplementedArticleServiceServer
	a  *biz.ArticleCase
	ac *biz.ArticleCategoryCase
}

func NewArticleService(a *biz.ArticleCase) *ArticleService {
	return &ArticleService{
		a: a,
	}
}

func (a *ArticleService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.Article, error) {
	articleDO := &biz.Article{
		Title:   in.GetTitle(),
		Desc:    in.GetDesc(),
		Content: in.GetContent(),
		Url:     in.GetUrl(),
		State:   int(in.GetState()),
		UserId:  in.GetUserId(),
	}
	err := a.a.Create(ctx, articleDO)
	if err != nil {
		return nil, err
	}
	//articleCategoryDO := &biz.ArticleCategory{
	//	ArticleId:  articleDO.Id,
	//	CategoryId: in.GetCategoryId(),
	//}
	//err = article.ac.Create(ctx, articleCategoryDO)
	//if err != nil {
	//	return nil, err
	//}
	return types.ArticleToStruct(articleDO), nil
}
func (a *ArticleService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.Article, error) {
	articleDO, err := a.a.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return types.ArticleToStruct(articleDO), nil
}
func (a *ArticleService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.Article, error) {
	articleD0 := &biz.Article{
		Title:   in.GetTitle(),
		Desc:    in.GetDesc(),
		Content: in.GetContent(),
		Url:     in.GetUrl(),
		State:   int(in.GetState()),
		UserId:  in.GetUserId(),
	}
	err := a.a.Update(ctx, in.GetId(), articleD0)
	if err != nil {
		return nil, err
	}
	return types.ArticleToStruct(articleD0), nil
}
func (a *ArticleService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*emptypb.Empty, error) {
	_, err := a.a.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (a *ArticleService) ListArticle(ctx context.Context, in *v1.ListArticleRequest) (*v1.ListArticleResponse, error) {
	articleDOs, err := a.a.List(ctx)
	if err != nil {
		return nil, err
	}
	articles := make([]*v1.Article, 0)
	for _, articleDO := range articleDOs {
		articles = append(articles, types.ArticleToStruct(articleDO))
	}
	return &v1.ListArticleResponse{
		Articles: articles,
	}, nil
}
