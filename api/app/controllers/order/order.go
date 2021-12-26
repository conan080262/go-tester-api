package ordercontroller

import (
	"net/http"
	"time"

	"github.com/conan080262/go-tester-api/configs"
	"github.com/conan080262/go-tester-api/models"
	"github.com/conan080262/go-tester-api/utilitys"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOrder(c *gin.Context) {
	var input OrderJson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}
	id, _ := primitive.ObjectIDFromHex(input.UserId)
	order := models.Order{
		ID:            primitive.NewObjectID(),
		UserPhone:     input.UserPhone,
		UserId:        id,
		StatusCurrent: "รอชำระเงิน",
		OrderAddress:  input.OrderAddress,
		ProductList:   input.ProductList,
		PriceDetail:   input.PriceDetail,
		Payment: models.Payment{
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
	}

	collection := configs.GetMongoCollection("orders")
	_, err := collection.InsertOne(configs.Ctx, order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":  "สร้างสำเร็จ",
		"data": order,
	})
}

func FetchOrder(c *gin.Context) {
	pagelimit := utilitys.StingToInt64(c.Query("pagelimit"))
	pageindex := utilitys.StingToInt64(c.Query("pageindex"))

	input := filter{
		TimeStart: c.Query("timestart"),
		TimeEnd:   c.Query("timeend"),
		Status:    c.Query("status"),
		Sort:      utilitys.StingToInt64(c.Query("sort")),
	}

	collection := configs.GetMongoCollection("orders")
	opts := options.Find()

	var query bson.D
	query = append(query, bson.E{})
	tst, _ := time.Parse(time.RFC3339, input.TimeStart)
	ten, _ := time.Parse(time.RFC3339, input.TimeEnd)
	if len(input.TimeStart) > 0 && len(input.TimeEnd) > 0 && input.Status == "ชำระเงินแล้ว" {
		query = append(query, bson.E{"payment.updated_at", bson.M{
			"$gt": tst.Add(-time.Hour * 7),
			"$lt": ten.Add(-time.Hour * 7),
		},
		})
	}
	if len(input.Status) > 0 {
		query = append(query, bson.E{"statuscurrent", input.Status})
	}

	if input.Sort == 1 || input.Sort == -1 {

	} else {
		input.Sort = -1
	}

	count, _ := collection.CountDocuments(configs.Ctx, query)
	paginate := utilitys.Paginate(pageindex, pagelimit, count)

	opts.SetSort(bson.D{{"payment.updated_at", input.Sort}})
	// opts.SetProjection(bson.M{"detail": 0})
	opts.SetLimit(paginate.PageLimit)
	opts.SetSkip(paginate.PageSkip)
	cursor, err := collection.Find(configs.Ctx, query, opts)

	if err != nil {

	}

	var results []bson.M
	if err = cursor.All(configs.Ctx, &results); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  tst,
		"page": paginate,
		"data": results,
	})
}

func UserAddPayment(c *gin.Context) {
	var input AddPayment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.Orderid)
	collection := configs.GetMongoCollection("orders")

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.M{
				"payment.date":       input.Payment.Date,
				"payment.bank":       input.Payment.Bank,
				"payment.pricetotal": input.Payment.PriceTotal,
				"payment.pricepaid":  input.Payment.PricePaid,
				"payment.name":       input.Payment.Name,
				"payment.accnum":     input.Payment.Accnum,
				"payment.imgbill":    input.Payment.Imgbill,
				"payment.updated_at": time.Now().Local(),
				"statuscurrent":      "ชำระเงินแล้ว",
			}},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบความผิดพลาด",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func StaffCarry(c *gin.Context) {
	var input Idorder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.Orderid)
	collection := configs.GetMongoCollection("orders")

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.M{
				"trackid":       input.TrackID,
				"statuscurrent": "รอรับสินค้า",
			}},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบความผิดพลาด",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func FinishOrder(c *gin.Context) {
	var input Idorder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.Orderid)
	collection := configs.GetMongoCollection("orders")
	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.M{
				"statuscurrent": "ส่งสินค้าสำเร็จ",
			}},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบความผิดพลาด",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func CancelOrder(c *gin.Context) {
	var input Idorder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.Orderid)
	collection := configs.GetMongoCollection("orders")
	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.M{
				"statuscurrent": "ยกเลิกคำสั่งซื้อ",
			}},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบความผิดพลาด",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func MyOrder(c *gin.Context) {
	var input Idorder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	var order models.Order
	id, _ := primitive.ObjectIDFromHex(input.Orderid)
	collection := configs.GetMongoCollection("orders")
	result := collection.FindOne(configs.Ctx, bson.M{"_id": id}).Decode(&order)

	if result != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not found this user order",
			"msg":   "ไม่พบคำสั่งซื้อ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "แสดงสินค้า",
		"data": order,
	})
}
