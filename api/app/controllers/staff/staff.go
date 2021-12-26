package staffcontroller

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

	// myUser := c.MustGet("user")
	user := models.Staff{
		ID:        primitive.NewObjectID(),
		Name:      input.Name,
		Role:      input.Role,
		Email:     input.Email,
		Password:  utilitys.HashPassword(input.Password),
		Phone:     input.Phone,
		IDCard:    input.IDCard,
		Gender:    input.Gender,
		BirthDate: input.BirthDate,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	collection := configs.GetMongoCollection("staffs")
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

	coll := configs.GetMongoCollection("staffs")
	var user models.Staff
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

func UploadImg(c *gin.Context) {
	var userObj imgUpload
	if err := c.ShouldBind(&userObj); err != nil {
		c.String(http.StatusBadRequest, "bad request1")
		return
	}

	path := uuid.NewString() + userObj.Avatar.Filename
	err := c.SaveUploadedFile(userObj.Avatar, "stores/"+path)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request2")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"path":   "/go/api/v1/public/images/" + path,
		"data":   userObj,
	})
}
