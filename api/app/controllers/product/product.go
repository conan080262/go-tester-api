package productcontroller

import (
	"net/http"
	"time"

	"github.com/conan080262/go-tester-api/configs"
	"github.com/conan080262/go-tester-api/models"
	"github.com/conan080262/go-tester-api/utilitys"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertProduct(c *gin.Context) {
	var input Productjson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	mapAllsize := make(map[string]int)
	for _, v := range input.Allsize {
		mapAllsize[v.Size] = v.Price
	}

	detailslice := make([]models.Details, len(input.Detail))
	var urlimgmain string
	var urlimgsub string
	for i, v := range input.Detail {
		detailslice[i].Sku = uuid.NewString() + "-" + input.Name + "-" + v.Size + "-" + v.Color
		detailslice[i].Size = v.Size
		detailslice[i].Qty = v.Qty
		detailslice[i].SelledAmount = 0
		detailslice[i].Price = mapAllsize[v.Size]
		detailslice[i].Color = v.Color
		detailslice[i].Urlimgmain = v.Urlimgmain
		detailslice[i].Urlimgsub = v.Urlimgsub

		if i == 0 {
			urlimgmain = v.Urlimgmain
			urlimgsub = v.Urlimgsub
		}
	}

	product := models.Product{
		ID:         primitive.NewObjectID(),
		Name:       input.Name,
		Category:   input.Category,
		Gender:     input.Gender,
		Allsize:    input.Allsize,
		Comment:    input.Comment,
		Urlimgmain: urlimgmain,
		Urlimgsub:  urlimgsub,
		Detail:     detailslice,
		Delete:     false,
		UpdatedBy:  c.MustGet("user").(models.Staff).ID,
		CreatedAt:  time.Now().Local(),
		UpdatedAt:  time.Now().Local(),
	}

	collection := configs.GetMongoCollection("products")
	_, err := collection.InsertOne(configs.Ctx, product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":  "สร้างสำเร็จ",
		"data": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var input Productjson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	mapAllsize := make(map[string]int)
	for _, v := range input.Allsize {
		mapAllsize[v.Size] = v.Price
	}

	detailslice := make([]models.Details, len(input.Detail))
	var urlimgmain string
	var urlimgsub string
	for i, v := range input.Detail {
		detailslice[i].Sku = uuid.NewString() + "-" + input.Name + "-" + v.Size + "-" + v.Color
		detailslice[i].Size = v.Size
		detailslice[i].Qty = v.Qty
		detailslice[i].SelledAmount = 0
		detailslice[i].Price = mapAllsize[v.Size]
		detailslice[i].Color = v.Color
		detailslice[i].Urlimgmain = v.Urlimgmain
		detailslice[i].Urlimgsub = v.Urlimgsub

		if i == 0 {
			urlimgmain = v.Urlimgmain
			urlimgsub = v.Urlimgsub
		}
	}
	id, _ := primitive.ObjectIDFromHex(input.IDupdate)
	product := models.ProductUpdate{
		Name:       input.Name,
		Category:   input.Category,
		Gender:     input.Gender,
		Allsize:    input.Allsize,
		Comment:    input.Comment,
		Urlimgmain: urlimgmain,
		Urlimgsub:  urlimgsub,
		Detail:     detailslice,
		Delete:     false,
		UpdatedBy:  c.MustGet("user").(models.Staff).ID,
		UpdatedAt:  time.Now().Local(),
	}

	collection := configs.GetMongoCollection("products")

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", product},
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "อัพเดทสำเร็จ",
		"data": result.ModifiedCount,
	})
}

func DeleteProduct(c *gin.Context) {
	var input Idproduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(input.ID)
	collection := configs.GetMongoCollection("products")
	result, err := collection.DeleteOne(configs.Ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "พบ Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ลบสำเร็จ",
		"data": result.DeletedCount,
	})
}

func FetchProduct(c *gin.Context) {
	pagelimit := utilitys.StingToInt64(c.Query("pagelimit"))
	pageindex := utilitys.StingToInt64(c.Query("pageindex"))

	input := filter{
		Gender:   c.Query("gender"),
		Category: c.Query("category"),
		Size:     c.Query("size"),
		SortKey:  c.Query("sortkey"),
		Sort:     utilitys.StingToInt64(c.Query("sort")),
	}

	collection := configs.GetMongoCollection("products")
	opts := options.Find()

	var query bson.D
	query = append(query, bson.E{})
	if len(input.Gender) > 0 {
		query = append(query, bson.E{"gender", input.Gender})
	}
	if len(input.Category) > 0 {
		query = append(query, bson.E{"category", input.Category})
	}
	if len(input.Size) > 0 {
		query = append(query, bson.E{"allsize.size", input.Size})
	}

	if len(input.SortKey) > 0 && (input.Sort == 1 || input.Sort == -1) {

	} else {
		input.SortKey = "_id"
		input.Sort = -1
	}

	count, _ := collection.CountDocuments(configs.Ctx, query)
	paginate := utilitys.Paginate(pageindex, pagelimit, count)

	opts.SetSort(bson.D{{input.SortKey, input.Sort}})
	opts.SetProjection(bson.M{"detail": 0})
	opts.SetLimit(paginate.PageLimit)
	opts.SetSkip(paginate.PageSkip)
	cursor, err := collection.Find(configs.Ctx, query, opts)

	if err != nil {

	}

	var results []bson.M
	if err = cursor.All(configs.Ctx, &results); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "ค้นหาสำเร็จ",
		"page": paginate,
		"data": results,
	})

}

func GetProduct(c *gin.Context) {
	productId := c.Query("id")
	id, _ := primitive.ObjectIDFromHex(productId)
	collection := configs.GetMongoCollection("products")
	var product models.Product
	result := collection.FindOne(configs.Ctx, bson.M{"_id": id}).Decode(&product)
	if result != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not found this email",
			"msg":   "ไม่พบสินค้าชนิดนี้",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "แสดงสินค้า",
		"data": product,
	})
}

func AddToUserCart(c *gin.Context) {
	var input CartUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	input.Cart.Uuid = uuid.NewString()
	id, _ := primitive.ObjectIDFromHex(input.ID)
	collection := configs.GetMongoCollection("users")

	result, err := collection.UpdateOne(
		configs.Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$push", bson.M{"cartlist": input.Cart}},
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
