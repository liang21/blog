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
	article *biz.ArticleCase
}

func NewArticleService(article *biz.ArticleCase) *ArticleService {
	return &ArticleService{
		article: article,
	}
}

func (article *ArticleService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.Article, error) {
	articleDO := &biz.Article{
		Title:   in.GetTitle(),
		Desc:    in.GetDesc(),
		Content: in.GetContent(),
		Url:     in.GetUrl(),
		State:   int(in.GetState()),
		UserId:  in.GetUserId(),
	}
	err := article.article.Create(ctx, articleDO)
	if err != nil {
		return nil, err
	}
	return types.ArticleToStruct(articleDO), nil
}
func (article *ArticleService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.Article, error) {
	return nil, nil
}
func (article *ArticleService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.Article, error) {
	return nil, nil
}
func (article *ArticleService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (article *ArticleService) ListArticle(ctx context.Context, in *v1.ListArticleRequest) (*v1.ListArticleResponse, error) {
	return nil, nil
}
