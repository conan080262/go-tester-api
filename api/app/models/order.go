package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	UserPhone     string             `json:"userphone" bson:"userphone"`
	UserId        primitive.ObjectID `json:"userid" bson:"userid"`
	StatusCurrent string             `json:"statuscurrent" bson:"statuscurrent"`
	Payment       Payment            `json:"payment" bson:"payment"`
	OrderAddress  OrderAddress       `json:"orderaddress" bson:"orderaddress"`
	ProductList   []ProductList      `json:"productlist" bson:"productlist"`
	TrackID       string             `json:"trackid" bson:"trackid"`
	PriceDetail   PriceDetail        `json:"pricedetail" bson:"pricedetail"`
}

type PriceDetail struct {
	Price      int `json:"price" bson:"price"`
	Vat        int `json:"vat" bson:"vat"`
	Carry      int `json:"carry" bson:"carry"`
	TotalPrice int `json:"totalprice" bson:"totalprice"`
}

type OrderAddress struct {
	Uuid         string `json:"uuid" bson:"uuid"`
	Name         string `json:"name" bson:"name"`
	Phonerecieve string `json:"phonerecieve" bson:"phonerecieve"`
	Provine      string `json:"provine" bson:"provine"`
	District     string `json:"district" bson:"district"`
	Sub_district string `json:"sub_district" bson:"sub_district"`
	Zipcode      string `json:"zipcode" bson:"zipcode"`
	Addressnum   string `json:"addtessnum" bson:"addressnum"`
	Comment      string `json:"comment" bson:"comment"`
}

type Payment struct {
	Date       string    `json:"date" bson:"date"`
	Bank       string    `json:"bank" bson:"bank"`
	PriceTotal int       `json:"pricetotal" bson:"pricetotal"`
	PricePaid  int       `json:"pricepaid" bson:"pricepaid"`
	Name       string    `json:"name" bson:"name"`
	Accnum     string    `json:"accnum" bson:"accnum"`
	Imgbill    string    `json:"imgbill" bson:"imgbill"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type ProductList struct {
	ProductID  primitive.ObjectID `json:"productid" bson:"productid"`
	Name       string             `json:"name" bson:"name"`
	Sku        string             `json:"sku" bson:"sku"`
	Size       string             `json:"size" bson:"size"`
	Qty        int                `json:"qty" bson:"qty"`
	Price      int                `json:"price" bson:"price"`
	PriceTotal int                `json:"pricetotal" bson:"pricetotal"`
	Color      string             `json:"color" bson:"color"`
	Category   string             `json:"category" bson:"category"`
	Gender     string             `json:"gender" bson:"gender"`
	Comment    string             `json:"comment" bson:"comment"`
	Urlimgmain string             `json:"imgmain" bson:"imgmain"`
	Urlimgsub  string             `json:"imgsub" bson:"imgsub"`
}
