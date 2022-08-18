package main

import(
	"fmt"
	// "github.com/jinzhu/gorm"
	// // _"github.com/jinzhu/gorm/dialects/sqlite"
	// "os"
	// "bufio"
	// "strings"
	// // "learn_go/models"
	// "learn_go/rand"
	// "learn_go/hash"
)

func main(){
	fmt.Println("Hello")

	// db := &gorm.DB{}

	// us := UserService{
	// 	uv: userValidator{
	// 		ug: userGorm{
	// 			db: db
	// 		}
	// 	}
	// }
	// fmt.Println(rand.String(32))
	// fmt.Println(rand.RememberToken())

	// h := hash.NewHMAC("mustafa")
	// fmt.Println(h.Hash("Hash it"))
	// fmt.Println(h.Hash("Hashit"))
	// fmt.Println(h.Hash("Hash it"))

	// fmt.Println("Connecting...")

	// us, _ := models.NewUserService("../db/lenslocked_dev.db")

	// us.DestructiveReset()

	// var user models.User
	// user.Name = "Michael Scott"
	// user.Email = "michael@dms.com"

	// err := us.Create(&user)

	// if err != nil{
	// 	fmt.Println(err)
	// }

	// fmt.Println(us.ByID(1))

	// user.Email = "michael@scottpaperco.com"

	// err = us.Update(&user)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(us.ByID(1))

	// fmt.Println(us.ByEmail("michael@scottpaperco.com"))

	// us.Delete(1)
	
	// us.Close()


	// db, err := gorm.Open("sqlite3", "../db/lenslocked_dev.db")
	// if err != nil{
	// 	panic(err)
	// }

	// db.LogMode(true)
	// fmt.Println("Connected...")
	// db.AutoMigrate(&User{})

	// name, email := getInfo()
	// user := User{Name:name, Email:email}
	// fmt.Println(user)

	// if err = db.Create(&user).Error

	// err != nil{
	// 	panic(err)
	// }
}


// func getInfo() (name, email string){
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("What is your name?")
// 	name, _ = reader.ReadString('\n')
// 	fmt.Println("What is your email address?")
// 	email, _ = reader.ReadString('\n')
// 	email = strings.TrimSpace(email)
// 	name = strings.TrimSpace(name)
// 	return name, email
// }