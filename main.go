package main
import (
    "github.com/gorilla/mux"
    "net/http"
    "learn_go/views"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/html")
    must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/html")
    must(contactView.Render(w, nil))
}

func main(){
    r := mux.NewRouter()

    homeView = views.NewView("bootstrap", "views/home.gohtml")
    contactView = views.NewView("bootstrap", "views/contact.gohtml")
    
    r.HandleFunc("/", home)
    r.HandleFunc("/contact", contact)
    http.ListenAndServe(":3000", r)
}

func must(err error){
    if err != nil{
        panic(err)
    }
}
