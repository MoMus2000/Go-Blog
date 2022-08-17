package main
import (
    "github.com/gorilla/mux"
    "net/http"
    "learn_go/controllers"
    "learn_go/models"
)

func main(){
    r := mux.NewRouter()
    us, _ := models.NewUserService("./db/lenslocked_dev.db")

    us.DestructiveReset()

    userC := controllers.NewUser(us)
    staticC := controllers.NewStaticView()
    
    r.Handle("/", staticC.HomeView).Methods("GET")
    r.Handle("/contact", staticC.ContactView).Methods("GET")
    
    r.HandleFunc("/signup", userC.New).Methods("GET")
    r.HandleFunc("/signup", userC.Create).Methods("POST")

    r.Handle("/login", userC.LoginView).Methods("GET")
    r.HandleFunc("/login", userC.Login).Methods("POST")

    r.HandleFunc("/cookie", userC.CookieTest).Methods("GET")

    http.ListenAndServe(":3000", r)
}

func must(err error){
    if err != nil{
        panic(err)
    }
}
