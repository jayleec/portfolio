package controllers

import (
	"github.com/revel/revel"
	"portfolio/app/models"
)

type Projects struct {
	App
}

func (c Projects) Index() revel.Result {
	return c.Render()
}

func (c Projects) Project(title string) revel.Result {
	results, err := c.Txn.Select(models.Project{},
	`select * from project where Title = ? `, title)
	if err != nil {
		panic(err)
	}
	var projects []*models.Project
	for _, r := range results {
		p := r.(*models.Project)
		projects = append(projects, p)
	}
	return c.Render(projects)
}


func (c Projects) loadById(id int) *models.Project {
	p, err := c.Txn.Get(models.Project{}, id)
	if err != nil{
		panic(err)
	}
	if p == nil{
		return nil
	}
	return p.(*models.Project)
}