package models

import (
	"context"
	"time"

	"github.com/conan080262/go-tester-api/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Email       string             `json:"email" bson:"email"`
	Nickname    string             `json:"nickname" bson:"nickname"`
	Password    string             `json:"-" bson:"password"`
	Gender      string             `json:"gender" bson:"gender"`
	BirthDate   string             `json:"birthdate" bson:"birthdate"`
	Phone       string             `json:"phone" bson:"phone"`
	Role        string             `json:"role" bson:"role"`
	AddressBook []AddressBooks     `json:"addressbook" bson:"addressbook"`
	Cartlist    []Cartlist         `json:"cartlist" bson:"cartlist"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type AddressBooks struct {
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

type Cartlist struct {
	Uuid       string `json:"uuid" bson:"uuid"`
	ProductID  string `json:"productid" bson:"productid"`
	Name       string `json:"name" bson:"name"`
	Sku        string `json:"sku" bson:"sku"`
	Size       string `json:"size" bson:"size"`
	Qty        int    `json:"qty" bson:"qty"`
	Price      int    `json:"price" bson:"price"`
	PriceTotal int    `json:"pricetotal" bson:"pricetotal"`
	Color      string `json:"color" bson:"color"`
	Category   string `json:"category" bson:"category"`
	Gender     string `json:"gender" bson:"gender"`
	Comment    string `json:"comment" bson:"comment"`
	Urlimgmain string `json:"imgmain" bson:"imgmain"`
	Urlimgsub  string `json:"imgsub" bson:"imgsub"`
}

func CreateUserManyIndexUnique(ctx context.Context) {
	coll := configs.GetMongoCollection("users")
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{"email", 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, _ = coll.Indexes().CreateMany(ctx, models, opts)
}
