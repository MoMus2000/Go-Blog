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

type SignupForm struct{
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){
	fmt.Println("HOLAAA")
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request){
	var form SignupForm

	err := parseRequestForm(r, &form)

	if err != nil{
		panic(err)
	}
	
	fmt.Println(form)
}