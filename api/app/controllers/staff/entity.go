package staffcontroller

import "mime/multipart"

type InputRegister struct {
	Name      string `json:"name" bson:"name"`
	Role      string `json:"role" bson:"role"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Phone     string `json:"phone" bson:"phone"`
	IDCard    string `json:"idcard" bson:"idcard"`
	Gender    string `json:"gender" bson:"gender"`
	BirthDate string `json:"birthdate" bson:"birthdate"`
}

type LoginJson struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type imgUpload struct {
	Avatar *multipart.FileHeader `form:"avatar"`
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
