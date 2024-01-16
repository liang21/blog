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

func CategoryToStruct(category *biz.Category) *v1.Category {
	return &v1.Category{
		Id:       category.Id,
		Name:     category.Name,
		Desc:     category.Desc,
		CreateAt: timestamppb.New(category.CreateAt),
		UpdateAt: timestamppb.New(category.UpdateAt),
	}
}

func TagToStruct(tag *biz.Tag) *v1.Tag {
	return &v1.Tag{
		Id:       tag.Id,
		Name:     tag.Name,
		Desc:     tag.Desc,
		CreateAt: timestamppb.New(tag.CreateAt),
		UpdateAt: timestamppb.New(tag.UpdateAt),
	}
}

func ArticleCategoryToStruct(articleCategory *biz.ArticleCategory) *v1.ArticleCategory {
	return &v1.ArticleCategory{
		Id:         articleCategory.Id,
		ArticleId:  articleCategory.ArticleId,
		CategoryId: articleCategory.CategoryId,
		CreateAt:   timestamppb.New(articleCategory.CreateAt),
		UpdateAt:   timestamppb.New(articleCategory.UpdateAt),
	}
}
