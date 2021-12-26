package usercontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/conan080262/go-tester-api/configs"
	"github.com/conan080262/go-tester-api/models"
	"github.com/conan080262/go-tester-api/utilitys"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	if !input.CheckRequired() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input not correct",
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}
	user := models.User{
		ID:          primitive.NewObjectID(),
		Email:       input.Email,
		Nickname:    input.Nickname,
		Password:    utilitys.HashPassword(input.Password),
		Gender:      input.Gender,
		BirthDate:   input.BirthDate,
		Phone:       input.Phone,
		Role:        "1111",
		Cartlist:    make([]models.Cartlist, 0),
		AddressBook: make([]models.AddressBooks, 0),
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	}

	collection := configs.GetMongoCollection("users")
	_, err := collection.InsertOne(configs.Ctx, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "มีอีเมลนี้ในระบบแล้ว",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "สร้างสำเร็จ",
		"data": user,
	})
}

func Login(c *gin.Context) {
	var input LoginJson

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	coll := configs.GetMongoCollection("users")
	var user models.User
	result := coll.FindOne(configs.Ctx, bson.M{"email": input.Email}).Decode(&user)

	if result != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not found this email",
			"msg":   "ไม่พบอีเมลนี้ในระบบ",
		})
		return
	}

	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "password not correct",
			"msg":   "รหัสผ่านไม่ถูกต้อง",
		})
		return
	}

	//สร้าง token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Email,
		"id":      user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	token, _ := claims.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusCreated, gin.H{
		"msg":          "เข้าสู้ระบบสำเร็จ",
		"access_token": token,
		"data":         user,
	})

}

func AddToAddressBooks(c *gin.Context) {
	var input MyAddBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}
	input.AddressBook.Uuid = uuid.NewString()
	id, _ := primitive.ObjectIDFromHex(input.ID)
	collection := configs.GetMongoCollection("users")

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$push", bson.M{"addressbook": input.AddressBook}},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบข้อผิดพลาด",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func PullUserAddr(c *gin.Context) {
	var input MyID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.ID)
	collection := configs.GetMongoCollection("users")

	change := bson.M{"$pull": bson.M{"addressbook": bson.M{"uuid": input.Uuid}}}

	result, err := collection.UpdateMany(
		configs.Ctx,
		bson.M{"_id": id},
		change,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ลบสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func PullUserCart(c *gin.Context) {
	var input MyID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.ID)
	collection := configs.GetMongoCollection("users")

	change := bson.M{"$pull": bson.M{"cartlist": bson.M{"uuid": input.Uuid}}}

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		change,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ลบสำเร็จ",
		"data": result.ModifiedCount,
	})
}
