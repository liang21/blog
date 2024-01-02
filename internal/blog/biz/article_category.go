package biz

import "context"

type ArticleCategory struct {
	Id         int64  `json:"id"`
	ArticleId  string `json:"article_id"`
	CategoryId string `json:"category_id"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
}

func (ArticleCategory) TableName() string {
	return "article_category"
}

type ArticleCategoryRepo interface {
	// db
	ListArticleCategory(ctx context.Context) ([]*ArticleCategory, error)
	GetArticleCategory(ctx context.Context, id int64) (*ArticleCategory, error)
	CreateArticleCategory(ctx context.Context, articleCategory *ArticleCategory) error
	UpdateArticleCategory(ctx context.Context, id int64, articleCategory *ArticleCategory) error
	DeleteArticleCategory(ctx context.Context, id int64) error
}
