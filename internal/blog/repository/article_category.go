package repository

import (
	"context"
	"errors"
	"github.com/liang21/blog/internal/blog/biz"
	"time"
	"xorm.io/xorm"
)

type acRepo struct {
	db *xorm.Engine
}

func (a *acRepo) ListArticleCategory(ctx context.Context) ([]*biz.ArticleCategory, error) {
	articleCategorys := make([]*biz.ArticleCategory, 0)
	err := a.db.Find(&articleCategorys)
	if err != nil {
		return nil, err
	}
	return articleCategorys, nil
}

func (a *acRepo) GetArticleCategory(ctx context.Context, articleId int64) (*biz.ArticleCategory, error) {
	articleCategory := &biz.ArticleCategory{CategoryId: articleId}
	ok, err := a.db.Get(articleCategory)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("articleCategory not found")
	}
	return articleCategory, nil
}

func (a *acRepo) CreateArticleCategory(ctx context.Context, articleCategory *biz.ArticleCategory) error {
	now := time.Now()
	articleCategory.CreateAt = now
	articleCategory.UpdateAt = now
	result, err := a.db.Insert(articleCategory)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (a *acRepo) UpdateArticleCategory(ctx context.Context, id int64, articleCategory *biz.ArticleCategory) error {
	articleCategory.UpdateAt = time.Now()
	result, err := a.db.ID(id).Update(articleCategory)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (a *acRepo) DeleteArticleCategory(ctx context.Context, id int64) error {
	result, err := a.db.ID(id).Delete(&biz.ArticleCategory{})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func NewArticleCategoryRepo(db *xorm.Engine) biz.ArticleCategoryRepo {
	return &acRepo{
		db: db,
	}
}
