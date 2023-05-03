package dao

import (
	"github.com/MyBlog/internal/model"
	"github.com/MyBlog/pkg/app"
)

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title string, state uint8, createdBy string) error {
	article := model.Article{
		Title: title,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
