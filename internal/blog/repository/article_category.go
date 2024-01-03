package repository

import (
	"context"
	"github.com/liang21/blog/internal/blog/biz"
	"xorm.io/xorm"
)

type acRepo struct {
	db *xorm.Engine
}

func (a acRepo) ListArticleCategory(ctx context.Context) ([]*biz.ArticleCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a acRepo) GetArticleCategory(ctx context.Context, id int64) (*biz.ArticleCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a acRepo) CreateArticleCategory(ctx context.Context, articleCategory *biz.ArticleCategory) error {
	//TODO implement me
	panic("implement me")
}

func (a acRepo) UpdateArticleCategory(ctx context.Context, id int64, articleCategory *biz.ArticleCategory) error {
	//TODO implement me
	panic("implement me")
}

func (a acRepo) DeleteArticleCategory(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewArticleCategoryRepo(db *xorm.Engine) biz.ArticleCategoryRepo {
	return &acRepo{
		db: db,
	}
}
