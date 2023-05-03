package service

import (
	"github.com/MyBlog/internal/model"
	"github.com/MyBlog/pkg/app"
)

type CountArticleRequest struct {
	Title string `form::"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form::"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title     string `form:"title" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"min=2,max=100"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.State, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
