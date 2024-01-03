package repository

import (
	"context"
	"errors"
	"github.com/liang21/blog/internal/blog/biz"
	"time"
	"xorm.io/xorm"
)

type categoryRepo struct {
	db *xorm.Engine
}

func NewCategoryRepo(db *xorm.Engine) biz.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) ListCategory(ctx context.Context) ([]*biz.Category, error) {
	categorys := make([]*biz.Category, 0)
	err := c.db.Find(&categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (c *categoryRepo) GetCategory(ctx context.Context, id int64) (*biz.Category, error) {
	category := &biz.Category{Id: id}
	ok, err := c.db.Get(category)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (c *categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) error {
	now := time.Now().Unix()
	category.CreateAt = now
	category.UpdateAt = now
	result, err := c.db.Insert(category)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (c *categoryRepo) UpdateCategory(ctx context.Context, id int64, category *biz.Category) error {
	category.UpdateAt = time.Now().Unix()
	result, err := c.db.ID(id).Update(category)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (c *categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	result, err := c.db.ID(id).Delete(&biz.Category{})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("delete failed")
	}
	return nil
}
