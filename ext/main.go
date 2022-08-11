package main

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"bufio"
	"strings"
)

type User struct{
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}

func main(){
	fmt.Println("Connecting...")
	db, err := gorm.Open("sqlite3", "../db/lenslocked_dev.db")
	if err != nil{
		panic(err)
	}

	db.LogMode(true)
	fmt.Println("Connected...")
	db.AutoMigrate(&User{})

	name, email := getInfo()
	user := User{Name:name, Email:email}
	fmt.Println(user)

	if err = db.Create(&user).Error

	err != nil{
		panic(err)
	}
}


func getInfo() (name, email string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email address?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)
	return name, email
}