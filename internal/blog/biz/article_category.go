package biz

import (
	"context"
	"time"
)

type ArticleCategory struct {
	Id         int64     `json:"id"`
	ArticleId  int64     `json:"article_id"`
	CategoryId int64     `json:"category_id"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}

func (ArticleCategory) TableName() string {
	return "article_category"
}

type ArticleCategoryRepo interface {
	// db
	ListArticleCategory(ctx context.Context) ([]*ArticleCategory, error)
	GetArticleCategory(ctx context.Context, articleId int64) (*ArticleCategory, error)
	CreateArticleCategory(ctx context.Context, articleCategory *ArticleCategory) error
	UpdateArticleCategory(ctx context.Context, id int64, articleCategory *ArticleCategory) error
	DeleteArticleCategory(ctx context.Context, id int64) error
}

type ArticleCategoryCase struct {
	repo ArticleCategoryRepo
}

func NewArticleCategoryCase(repo ArticleCategoryRepo) *ArticleCategoryCase {
	return &ArticleCategoryCase{repo: repo}
}

func (a *ArticleCategoryCase) List(ctx context.Context) ([]*ArticleCategory, error) {
	return a.repo.ListArticleCategory(ctx)
}

func (a *ArticleCategoryCase) Get(ctx context.Context, articleId int64) (*ArticleCategory, error) {
	return a.repo.GetArticleCategory(ctx, articleId)
}

func (a *ArticleCategoryCase) Create(ctx context.Context, articleCategory *ArticleCategory) error {
	return a.repo.CreateArticleCategory(ctx, articleCategory)
}

func (a *ArticleCategoryCase) Update(ctx context.Context, id int64, articleCategory *ArticleCategory) error {
	return a.repo.UpdateArticleCategory(ctx, id, articleCategory)
}

func (a *ArticleCategoryCase) Delete(ctx context.Context, id int64) error {
	return a.repo.DeleteArticleCategory(ctx, id)
}
