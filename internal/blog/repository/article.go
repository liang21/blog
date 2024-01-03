package repository

import (
	"context"
	"errors"
	"github.com/liang21/blog/internal/blog/biz"
	"time"
	"xorm.io/xorm"
)

type articleRepo struct {
	db *xorm.Engine
}

func NewArticleRepo(db *xorm.Engine) biz.ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

func (a *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	articles := make([]*biz.Article, 0)
	err := a.db.Find(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	article := &biz.Article{Id: id}
	ok, err := a.db.Get(article)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (a *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	now := time.Now().Unix()
	article.CreateAt = now
	article.UpdateAt = now
	result, err := a.db.Insert(article)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (a *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	article.UpdateAt = time.Now().Unix()
	result, err := a.db.ID(id).Update(article)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (a *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	article := &biz.Article{Id: id}
	result, err := a.db.Delete(article)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (a *articleRepo) GetArticleLike(ctx context.Context, id int64) (rv int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleRepo) IncArticleLike(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
