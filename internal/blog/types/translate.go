package types

import (
	v1 "github.com/liang21/blog/api/blog/v1"
	"github.com/liang21/blog/internal/blog/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ArticleToStruct(article *biz.Article) *v1.Article {
	return &v1.Article{
		Id:       article.Id,
		Title:    article.Title,
		Desc:     article.Desc,
		Content:  article.Content,
		Url:      article.Url,
		State:    int64(int32(article.State)),
		UserId:   article.UserId,
		CreateAt: timestamppb.New(article.CreateAt),
		UpdateAt: timestamppb.New(article.UpdateAt),
	}
}
