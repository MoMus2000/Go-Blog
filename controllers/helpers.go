package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
	"fmt"
)

func parseRequestForm(r *http.Request, form interface{}) error{
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	fmt.Println("Parsing the form !")
	
	decoder := schema.NewDecoder()

	err = decoder.Decode(form, r.PostForm)

	if err != nil{
		panic(err)
	}

	fmt.Println("Parsed the form !")

	return nil
}