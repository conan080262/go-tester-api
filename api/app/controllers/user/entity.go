package usercontroller

import "github.com/conan080262/go-tester-api/models"

type InputRegister struct {
	Email     string `json:"email" bson:"email"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Password  string `json:"password" bson:"password"`
	Gender    string `json:"gender" bson:"gender"`
	BirthDate string `json:"birthdate" bson:"birthdate"`
	Phone     string `json:"phone" bson:"phone"`
}

type LoginJson struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type MyAddBook struct {
	ID          string              `json:"userid" bson:"userid"`
	AddressBook models.AddressBooks `json:"addressbook" bson:"addressbook"`
}

type MyID struct {
	ID   string `json:"userid" bson:"userid"`
	Uuid string `json:"uuid" bson:"uuid"`
}

func (input *InputRegister) CheckRequired() bool {
	checked := true
	if input.Email == "" {
		checked = false
	} else if input.Password == "" {
		checked = false
	}
	return checked
}
