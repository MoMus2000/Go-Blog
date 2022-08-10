package controllers

import (
	"learn_go/views"
	"net/http"
	"fmt"
	"github.com/gorilla/schema"
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
	
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello we are here")

	var form SignupForm
	
	decoder := schema.NewDecoder()

	err = decoder.Decode(&form, r.PostForm)

	if err != nil{
		panic(err)
	}

	fmt.Println(form)
}