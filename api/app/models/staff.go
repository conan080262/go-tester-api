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

type Staff struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Role      string             `json:"role" bson:"role"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	Phone     string             `json:"phone" bson:"phone"`
	IDCard    string             `json:"idcard" bson:"idcard"`
	Gender    string             `json:"gender" bson:"gender"`
	BirthDate string             `json:"birthdate" bson:"birthdate"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func CreateStaffManyIndexUnique(ctx context.Context) {
	coll := configs.GetMongoCollection("staffs")
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{"email", 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, _ = coll.Indexes().CreateMany(ctx, models, opts)
}
