package repository

import (
	"context"
	"github.com/liang21/blog/internal/blog/biz"
)

type categoryRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func (c *categoryRepo) ListCategory(ctx context.Context) ([]*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepo) GetCategory(ctx context.Context, id int64) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepo) UpdateCategory(ctx context.Context, id int64, category *biz.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewCategoryRepo(db *xorm.Engine, rdb *redis.Client) biz.CategoryRepo {
	return &categoryRepo{
		db:  db,
		rdb: rdb,
	}
}
