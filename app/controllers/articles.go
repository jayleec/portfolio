package controllers

import (
	"github.com/revel/revel"
	"portfolio/app/models"
)

type Articles struct {
	App
}

func (c Articles) Index() revel.Result {
	return c.Render()
}

func (c Articles) Article(address string) revel.Result {
	results, err := c.Txn.Select(models.Article{},
		`select * from article where Address = ? `, address)
	if err != nil {
		panic(err)
	}
	var articles []*models.Article
	for _, r := range results {
		a := r.(*models.Article)
		articles = append(articles, a)
	}
	return c.Render(articles)
}