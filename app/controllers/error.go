package controllers

import (
	"github.com/revel/revel"
)

type ErrorPages struct {
	*revel.Controller
}

func (e ErrorPages) Page404() revel.Result {
	return e.Render()
}

func (e ErrorPages) Page500() revel.Result {
	return e.Render()
}
