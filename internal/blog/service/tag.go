package service

import (
	"context"
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
	"github.com/liang21/blog/internal/blog/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TagService struct {
	v1.UnimplementedTagServiceServer
	t *biz.TagCase
}

func NewTagService(t *biz.TagCase) *TagService {
	return &TagService{
		t: t,
	}
}

func (t *TagService) CreateTag(ctx context.Context, in *v1.CreateTagRequest) (*v1.Tag, error) {
	tagDO := &biz.Tag{
		Name: in.GetName(),
		Desc: in.GetDesc(),
	}
	err := t.t.Create(ctx, tagDO)
	if err != nil {
		return nil, err
	}
	return types.TagToStruct(tagDO), nil
}
func (t *TagService) GetTag(ctx context.Context, in *v1.GetTagRequest) (*v1.Tag, error) {
	tagDO, err := t.t.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return types.TagToStruct(tagDO), nil
}
func (t *TagService) UpdateTag(ctx context.Context, in *v1.UpdateTagRequest) (*v1.Tag, error) {
	tagDO := &biz.Tag{
		Name: in.GetName(),
		Desc: in.GetDesc(),
	}
	err := t.t.Update(ctx, in.GetId(), tagDO)
	if err != nil {
		return nil, err
	}
	return types.TagToStruct(tagDO), nil
}
func (t *TagService) DeleteTag(ctx context.Context, in *v1.DeleteTagRequest) (*emptypb.Empty, error) {
	err := t.t.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (t *TagService) ListTag(ctx context.Context, in *v1.ListTagRequest) (*v1.ListTagResponse, error) {
	tagDOs, err := t.t.List(ctx)
	if err != nil {
		return nil, err
	}
	tags := make([]*v1.Tag, 0)
	for _, tagDO := range tagDOs {
		tags = append(tags, types.TagToStruct(tagDO))
	}
	return &v1.ListTagResponse{Tags: tags}, nil
}
