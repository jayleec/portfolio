package controllers

import "github.com/revel/revel"

type App struct {
	GorpController
}

func (c App) Index() revel.Result{
	//return c.Render()
	return c.Redirect(Projects.Index)
}


func (c App) Articles() revel.Result {
	return c.Render()
}

func (c App) Awards() revel.Result {
	return c.Render()
}

func (c App) Golang() revel.Result {
	return c.Render()
}