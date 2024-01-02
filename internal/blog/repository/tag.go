package repository

import (
	"context"
	"github.com/liang21/blog/internal/blog/biz"
)

type tagRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewTagRepo(db *xorm.Engine, rdb *redis.Client) biz.TagRepo {
	return &tagRepo{
		db:  db,
		rdb: rdb,
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
