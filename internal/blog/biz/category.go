package biz

import "context"

type Category struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRepo interface {
	// db
	ListCategory(ctx context.Context) ([]*Category, error)
	GetCategory(ctx context.Context, id int64) (*Category, error)
	CreateCategory(ctx context.Context, category *Category) error
	UpdateCategory(ctx context.Context, id int64, category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}
