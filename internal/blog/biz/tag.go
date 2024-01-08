package biz

import (
	"context"
	"time"
)

type Tag struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (Tag) TableName() string {
	return "tag"
}

type TagRepo interface {
	// db
	ListTag(ctx context.Context) ([]*Tag, error)
	GetTag(ctx context.Context, id int64) (*Tag, error)
	CreateTag(ctx context.Context, tag *Tag) error
	UpdateTag(ctx context.Context, id int64, tag *Tag) error
	DeleteTag(ctx context.Context, id int64) error
}

type TagCase struct {
	repo TagRepo
}

func NewTagCase(repo TagRepo) *TagCase {
	return &TagCase{repo: repo}
}

func (t *TagCase) List(ctx context.Context) ([]*Tag, error) {
	return t.repo.ListTag(ctx)
}

func (t *TagCase) Get(ctx context.Context, id int64) (*Tag, error) {
	return t.repo.GetTag(ctx, id)
}

func (t *TagCase) Create(ctx context.Context, tag *Tag) error {
	return t.repo.CreateTag(ctx, tag)
}

func (t *TagCase) Update(ctx context.Context, id int64, tag *Tag) error {
	return t.repo.UpdateTag(ctx, id, tag)
}

func (t *TagCase) Delete(ctx context.Context, id int64) error {
	return t.repo.DeleteTag(ctx, id)
}
