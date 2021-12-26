package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Category   string             `json:"category" bson:"category"`
	Gender     string             `json:"gender" bson:"gender"`
	Allsize    []Allsize          `json:"allsize" bson:"allsize"`
	Comment    string             `json:"comment" bson:"comment"`
	Urlimgmain string             `json:"urlimgmain" bson:"urlimgmain"`
	Urlimgsub  string             `json:"urlimgsub" bson:"urlimgsub"`
	Detail     []Details          `json:"detail" bson:"detail"`
	Delete     bool               `json:"delete" bson:"delete"`
	UpdatedBy  primitive.ObjectID `json:"updatedby" bson:"updatedby"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type ProductUpdate struct {
	Name       string             `json:"name" bson:"name"`
	Category   string             `json:"category" bson:"category"`
	Gender     string             `json:"gender" bson:"gender"`
	Allsize    []Allsize          `json:"allsize" bson:"allsize"`
	Comment    string             `json:"comment" bson:"comment"`
	Urlimgmain string             `json:"urlimgmain" bson:"urlimgmain"`
	Urlimgsub  string             `json:"urlimgsub" bson:"urlimgsub"`
	Detail     []Details          `json:"detail" bson:"detail"`
	Delete     bool               `json:"delete" bson:"delete"`
	UpdatedBy  primitive.ObjectID `json:"updatedby" bson:"updatedby"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type Allsize struct {
	Size  string `json:"size" bson:"size"`
	Price int    `json:"price" bson:"price"`
}

type Details struct {
	Sku          string `json:"sku" bson:"sku"`
	Size         string `json:"size" bson:"size"`
	Qty          int    `json:"qty" bson:"qty"`
	SelledAmount int    `json:"selledamount" bson:"selledamount"`
	Price        int    `json:"price" bson:"price"`
	Color        string `json:"color" bson:"color"`
	Urlimgmain   string `json:"urlimgmain" bson:"urlimgmain"`
	Urlimgsub    string `json:"urlimgsub" bson:"urlimgsub"`
}
