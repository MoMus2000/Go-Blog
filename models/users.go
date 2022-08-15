package models


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotFound = errors.New("models: resource not found")
var InvalidId = errors.New("id provided was invalid")

const passwordPepper = "Salt&Peppa"

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
	pwBytes := []byte(u.Password + passwordPepper)
	hashedBytes, error := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)

	if error != nil{
		panic(error)
	}

	u.PasswordHash = string(hashedBytes)
	u.Password = ""

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

func (us *UserService) Authenticate(email string, password string) (*User, error){
	foundUser, error := us.ByEmail(email)

	if error != nil {
		return nil, error
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password + passwordPepper))

	if err != nil{
		return nil, err
	}

	return foundUser, nil
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


func (us *UserService) DestructiveReset() error{ 
	error := us.db.DropTableIfExists(&User{}).Error
	if error != nil {
		return error
	}
	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error{
	err := us.db.AutoMigrate(&User{}).Error
	if err != nil{
		return err
	}
	return nil
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
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not_null"`
}