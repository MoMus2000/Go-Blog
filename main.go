package main
import (
    "github.com/gorilla/mux"
    "net/http"
    "learn_go/views"
    "learn_go/controllers"
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
    userC := controllers.NewUser()
    
    r.HandleFunc("/", home).Methods("GET")
    r.HandleFunc("/contact", contact).Methods("GET")
    r.HandleFunc("/signup", userC.New).Methods("GET")
    r.HandleFunc("/signup", userC.Create).Methods("POST")

    http.ListenAndServe(":3000", r)
}

func must(err error){
    if err != nil{
        panic(err)
    }
}
