package repository

import (
	"context"
	"errors"
	"github.com/liang21/blog/internal/blog/biz"
	"time"
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
	tags := make([]*biz.Tag, 0)
	err := t.db.Find(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *tagRepo) GetTag(ctx context.Context, id int64) (*biz.Tag, error) {
	tag := &biz.Tag{Id: id}
	ok, err := t.db.Get(tag)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("tag not found")
	}
	return tag, nil
}

func (t *tagRepo) CreateTag(ctx context.Context, tag *biz.Tag) error {
	now := time.Now().Unix()
	tag.CreateAt = now
	tag.UpdateAt = now
	result, err := t.db.Insert(tag)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (t *tagRepo) UpdateTag(ctx context.Context, id int64, tag *biz.Tag) error {
	tag.UpdateAt = time.Now().Unix()
	result, err := t.db.ID(id).Update(tag)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (t *tagRepo) DeleteTag(ctx context.Context, id int64) error {
	tag := &biz.Tag{Id: id}
	result, err := t.db.Delete(tag)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("delete failed")
	}
	return nil
}
