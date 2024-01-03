package biz

import "context"

type Article struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	Url      string `json:"url"`
	State    int    `json:"state"`
	UserId   int64  `json:"user_id"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleLike(ctx context.Context, id int64) (rv int64, err error)
	IncArticleLike(ctx context.Context, id int64) error
}

type ArticleCase struct {
	repo ArticleRepo
}

func NewArticleCase(repo ArticleRepo) *ArticleCase {
	return &ArticleCase{repo: repo}
}

func (a *ArticleCase) List(ctx context.Context) ([]*Article, error) {
	return a.repo.ListArticle(ctx)
}

func (a *ArticleCase) Get(ctx context.Context, id int64) (*Article, error) {
	return a.repo.GetArticle(ctx, id)
}

func (a *ArticleCase) Create(ctx context.Context, article *Article) error {
	return a.repo.CreateArticle(ctx, article)
}

func (a *ArticleCase) Update(ctx context.Context, id int64, article *Article) error {
	return a.repo.UpdateArticle(ctx, id, article)
}

func (a *ArticleCase) Delete(ctx context.Context, id int64) error {
	return a.repo.DeleteArticle(ctx, id)
}
