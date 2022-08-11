package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

func getAlltemplateFiles() []string{
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil{
		panic(err)
	}
	return files
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request){
	err := v.Render(w, nil)
	if err != nil{
		panic(err)
	}
}

func (view *View) Render(w http.ResponseWriter, data interface{}) error{
	return view.Template.ExecuteTemplate(w, view.Layout, data)
}

func NewView(layout string, files ...string) *View {
	files = append(files, getAlltemplateFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template : t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}