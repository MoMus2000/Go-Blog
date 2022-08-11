package main
import (
    "github.com/gorilla/mux"
    "net/http"
    "learn_go/controllers"
)

func main(){
    r := mux.NewRouter()
    userC := controllers.NewUser()
    staticC := controllers.NewStaticView()
    
    r.Handle("/", staticC.HomeView).Methods("GET")
    r.Handle("/contact", staticC.ContactView).Methods("GET")
    r.HandleFunc("/signup", userC.New).Methods("GET")
    r.HandleFunc("/signup", userC.Create).Methods("POST")

    http.ListenAndServe(":3000", r)
}

func must(err error){
    if err != nil{
        panic(err)
    }
}
