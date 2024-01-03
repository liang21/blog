package repository

import (
	"context"
	"github.com/liang21/blog/internal/blog/biz"
	"xorm.io/xorm"
)

type tagRepo struct {
	db *xorm.Engine
}

func NewTagRepo(db *xorm.Engine) biz.TagRepo {
	return &tagRepo{
		db: db,
	}
}

var _ biz.TagRepo = (*tagRepo)(nil)

func (t *tagRepo) ListTag(ctx context.Context) ([]*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepo) GetTag(ctx context.Context, id int64) (*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepo) CreateTag(ctx context.Context, tag *biz.Tag) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepo) UpdateTag(ctx context.Context, id int64, tag *biz.Tag) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepo) DeleteTag(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
