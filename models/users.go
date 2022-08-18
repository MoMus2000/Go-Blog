package models


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"learn_go/hash"
	"learn_go/rand"
)

var ErrNotFound = errors.New("models: resource not found")
var InvalidId = errors.New("id provided was invalid")

const passwordPepper = "Salt&Peppa"
const secretKey = "secret-key"

type UserService struct{
	db *gorm.DB
	hmac hash.HMAC
}

type UserDB interface{
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRememberToken(token string) (*User, error)

	// Methods for altering users
	Create(user *User) (error)
	Update(user *User) (error)
	Delete(id uint) (error)

	Close() (error)

	AutoMigrate() error
	DestructiveReset() error
}


// compilation test to see if userGorm implements all the methods for the interface
var _ UserDB = &userGorm{}

type userGorm struct{
	db *gorm.DB
	hamc hash.HMAC
}

func NewUserService(connectionInfo string) (*UserService, error){
	db, err:= gorm.Open("sqlite3", connectionInfo)
	if err != nil{
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(secretKey)
	return &UserService{
		db: db,
		hmac : hmac,
	}, nil
}

func NewGormService(connectionInfo string) (*userGorm, error){
	db, err := gorm.Open("sqlite3", connectionInfo)
	if err != nil {
		return nil , error
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(secretKey)
	return &userGorm{
		db: db,
		hmac: hmac
	}, nil
}

func (ug *userGorm) Create(u *User) error{
	pwBytes := []byte(u.Password + passwordPepper)
	hashedBytes, error := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)

	if error != nil{
		panic(error)
	}

	u.PasswordHash = string(hashedBytes)
	u.Password = ""

	if u.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil{
			return err
		}
		u.Remember = token
	}

	u.RememberHash = ug.hmac.Hash(u.Remember)

	return ug.db.Create(u).Error
}

func (ug *userGorm) Update(u *User) error{
	if u.Remember != "" {
		u.RememberHash = ug.hmac.Hash(u.Remember)
	}
	return ug.db.Save(u).Error
}

func (ug *userGorm) ByEmail(email string) (*User, error){
	var user User

	db := ug.db.Where("email = ?", email)
	err := first(db, &user)

	return &user, err	
}

func (ug *userGorm) ByRememberToken(token string) (*User, error){
	var user User
	hashedToken := ug.hmac.Hash(token)
	db := ug.db.Where("remember_hash = ?", hashedToken)
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

func (ug *userGorm) Delete (id uint) error{
	if id == 0{
		return InvalidId
	}

	user := User{Model:gorm.Model{ID:id}}
	return ug.db.Delete(&user).Error
}

func first(db *gorm.DB, dst interface{}) error{
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}


func (ug *userGorm) DestructiveReset() error{ 
	error := ug.db.DropTableIfExists(&User{}).Error
	if error != nil {
		return error
	}
	return ug.AutoMigrate()
}

func (ug *userGorm) AutoMigrate() error{
	err := ug.db.AutoMigrate(&User{}).Error
	if err != nil{
		return err
	}
	return nil
}

// ByID will look up user by the id provided
// case 1 - user, nil
// case 2 - nil, ErrNotFound
// case 3 - nil, otherError 
func (ug *userGorm) ByID(id uint) (*User, error){
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) Close() error{
	return ug.db.Close()
}

type User struct{
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not_null"`
	Remember string `gorm:"-"`
	RememberHash string `gorm:"not_null;unique_index"`
}