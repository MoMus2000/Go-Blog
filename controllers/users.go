package controllers

import (
	"learn_go/views"
	"net/http"
	"fmt"
	"learn_go/models"
	"learn_go/rand"
)

func NewUser(us *models.UserService) *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/signup.gohtml"),
		LoginView: views.NewView("bootstrap", "views/users/login.gohtml"),
		us: us,
	}
}

type Users struct{
	NewView *views.View
	LoginView *views.View
	us *models.UserService
}


type SignupForm struct{
	Name string `schema:"name"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct{
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request){
	var form SignupForm

	err := parseRequestForm(r, &form)

	if err != nil{
		panic(err)
	}

	user := models.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}

	err = u.us.Create(&user)

	if err != nil{
		panic(err)
	}
	
	fmt.Println(form)

	err = u.SignIn(w, &user)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/cookie", http.StatusFound)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request){
	var form LoginForm
	err := parseRequestForm(r, &form)

	if err != nil{
		panic(err)
	}

	user, error := u.us.Authenticate(form.Email, form.Password)

	if error != nil{
		panic(error)
	}

	err = u.SignIn(w, user)

	if err != nil{
		panic(err)
	}

	http.Redirect(w, r, "/cookie", http.StatusFound)
} 

func (u *Users) SignIn(w http.ResponseWriter, user *models.User) error{
	if user.Remember == ""{
		token, err := rand.RememberToken()
		if err != nil{
 			return err
		}
		user.Remember = token

		err = u.us.Update(user)
	}
	cookie := http.Cookie{ 
		Name: "remember_token",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)

	return nil
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("remember_token")

	if err != nil{
		panic(err)
	}

	fmt.Println(cookie)

	fmt.Fprintln(w, cookie)

	user, err := u.us.ByRememberToken(cookie.Value)

	if err != nil{
		fmt.Fprintln(w, user)
		fmt.Fprintln(w, "Resource not found")
	}

	fmt.Println(w, user)
}