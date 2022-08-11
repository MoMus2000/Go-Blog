package controllers

import (
	"learn_go/views"
)

type Static struct{
	HomeView *views.View
	ContactView *views.View
}

func NewStaticView() *Static{
	return &Static{
		HomeView: views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}