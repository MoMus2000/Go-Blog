package models


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"errors"
)

var ErrNotFound = errors.New("models: resource not found")

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
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err{
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (us *UserService) Close() error{
	return us.db.Close()
}


type User struct{
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}