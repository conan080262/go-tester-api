package ordercontroller

import (
	"github.com/conan080262/go-tester-api/models"
)

type OrderJson struct {
	UserPhone    string               `json:"userphone" bson:"userphone"`
	UserId       string               `json:"userid" bson:"userid"`
	OrderAddress models.OrderAddress  `json:"orderaddress" bson:"orderaddress"`
	PriceDetail  models.PriceDetail   `json:"pricedetail" bson:"pricedetail"`
	ProductList  []models.ProductList `json:"productlist" bson:"productlist"`
}

type filter struct {
	TimeStart string `json:"timestart" bson:"timestart"`
	TimeEnd   string `json:"timeend" bson:"timeend"`
	Status    string `json:"status" bson:"status"`
	Sort      int64  `json:"sort" bson:"sort"`
}

type Idorder struct {
	Orderid string `json:"orderid" bson:"orderid"`
	TrackID string `json:"trackid" bson:"trackid"`
}

type AddPayment struct {
	Orderid string         `json:"orderid" bson:"orderid"`
	Payment models.Payment `json:"payment" bson:"payment"`
}
