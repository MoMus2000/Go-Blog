package controllers

import (
	"learn_go/views"
	"net/http"
	"fmt"
)

func NewUser() *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/signup.gohtml"),
	}
}

type Users struct{
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){
	fmt.Println("HOLAAA")
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello we are here")
	fmt.Fprintln(w, "Consider that the user has been created")
}