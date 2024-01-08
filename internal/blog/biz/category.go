package biz

import (
	"context"
	"time"
)

type Category struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
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

type CategoryCase struct {
	repo CategoryRepo
}

func NewCategoryCase(repo CategoryRepo) *CategoryCase {
	return &CategoryCase{repo: repo}
}

func (c *CategoryCase) List(ctx context.Context) ([]*Category, error) {
	return c.repo.ListCategory(ctx)
}

func (c *CategoryCase) Get(ctx context.Context, id int64) (*Category, error) {
	return c.repo.GetCategory(ctx, id)
}

func (c *CategoryCase) Create(ctx context.Context, category *Category) error {
	return c.repo.CreateCategory(ctx, category)
}

func (c *CategoryCase) Update(ctx context.Context, id int64, category *Category) error {
	return c.repo.UpdateCategory(ctx, id, category)
}

func (c *CategoryCase) Delete(ctx context.Context, id int64) error {
	return c.repo.DeleteCategory(ctx, id)
}
