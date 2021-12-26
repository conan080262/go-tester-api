package productcontroller

import "github.com/conan080262/go-tester-api/models"

type Productjson struct {
	IDupdate string           `json:"idupdate" bson:"idupdate"`
	Name     string           `json:"name" bson:"name"`
	Category string           `json:"category" bson:"category"`
	Gender   string           `json:"gender" bson:"gender"`
	Comment  string           `json:"comment" bson:"comment"`
	Allsize  []models.Allsize `json:"allsize" bson:"allsize"`
	Detail   []Detail         `json:"detail" bson:"detail"`
}

type filter struct {
	Gender   string
	Category string
	Size     string
	SortKey  string
	Sort     int64
}

type Idproduct struct {
	ID string `json:"id" bson:"id"`
}

type CartUser struct {
	ID   string          `json:"userid" bson:"userid"`
	Cart models.Cartlist `json:"cart" bson:"cart"`
}

type Detail struct {
	Size       string `json:"size" bson:"size"`
	Color      string `json:"color" bson:"color"`
	Urlimgmain string `json:"urlimgmain" bson:"urlimgmain"`
	Urlimgsub  string `json:"urlimgsub" bson:"urlimgsub"`
	Qty        int    `json:"qty" bson:"qty"`
}
