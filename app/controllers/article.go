package controllers

import "github.com/revel/revel"

type Articles struct {
	App
}

func (c Articles) Index() revel.Result {
	return c.Render()
}

func (c Articles) Article(title string) revel.Result {
	return c.Render()
}