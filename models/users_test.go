package models

import (
	"testing"
)

func testingUserService() (*UserService, error){
	us, err := NewUserService("../db/lenslocked_dev")
	if err != nil{
		return nil, err
	}

	us.db.LogMode(false)

	us.DestructiveReset()

	return us, nil
}

func TestCreateUser(t *testing.T){
	us, err := testingUserService()
	if err != nil{
		t.Fatal(err)
	}

	user := User{
		Name: "Michael Scott",
		Email: "mm@dm.com",
	}

	err = us.Create(&user)

	if err != nil{
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0")
	}

}