package biz

import "context"

type Tag struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
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
