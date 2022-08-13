package models


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"errors"
)

var ErrNotFound = errors.New("models: resource not found")
var InvalidId = errors.New("id provided was invalid")

type UserService struct{
	db *gorm.DB
}


func NewUserService(connectionInfo string) (*UserService, error){
	db, err:= gorm.Open("sqlite3", connectionInfo)
	if err != nil{
		return nil, err
	}
	db.LogMode(true)
	return &UserService{db: db}, nil
}

func (us *UserService) Create(u *User) error{
	return us.db.Create(u).Error
}

func (us *UserService) Update(u *User) error{
	return us.db.Save(u).Error
}

func (us *UserService) ByEmail(email string) (*User, error){
	var user User

	db := us.db.Where("email = ?", email)
	err := first(db, &user)

	return &user, err	
}

func (us *UserService) Delete (id uint) error{
	if id == 0{
		return InvalidId
	}

	user := User{Model:gorm.Model{ID:id}}
	return us.db.Delete(&user).Error
}

func first(db *gorm.DB, dst interface{}) error{
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}


func (us *UserService) DestructiveReset(){
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}


// ByID will look up user by the id provided
// case 1 - user, nil
// case 2 - nil, ErrNotFound
// case 3 - nil, otherError 
func (us *UserService) ByID(id uint) (*User, error){
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) Close() error{
	return us.db.Close()
}


type User struct{
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}