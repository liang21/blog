package repository

import (
	"context"
	"github.com/liang21/blog/internal/blog/biz"
	"xorm.io/xorm"
)

type articleRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewArticleRepo(db *xorm.Engine, rdb *redis.Client) biz.ArticleRepo {
	return &articleRepo{
		db:  db,
		rdb: rdb,
	}
}

func (a *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) GetArticleLike(ctx context.Context, id int64) (rv int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) IncArticleLike(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
