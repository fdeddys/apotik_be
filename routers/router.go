package routers

import (
	"distribution-system-be/controllers"
	"distribution-system-be/models"
	dto "distribution-system-be/models/dto"
	"distribution-system-be/utils/security"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	kons "distribution-system-be/constants"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	fmt.Println(gin.IsDebugging())
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	UserController := new(controllers.UserController)
	CustomerController := new(controllers.CustomerController)
	SupplierController := new(controllers.SupplierController)
	OrderController := new(controllers.OrderController)
	OrderDetailController := new(controllers.OrderDetailController)
	DashboardController := new(controllers.DashboardController)

	api := r.Group("/api/user")
	api.POST("/filter/page/:page/count/:count", UserController.GetUser)
	api.POST("/", UserController.SaveDataUser)
	api.PUT("/", UserController.UpdateUser)

	AuthController := new(controllers.AuthController)
	api = r.Group("/auth")
	api.POST("/login", AuthController.Login)
	api.GET("/get_cur_user", AuthController.GetCurrPass)
	api.POST("/change_pass", AuthController.ChangePass)

	api = r.Group("/api/customer")
	api.POST("/page/:page/count/:count", CustomerController.FilterDataCustomer)
	api.POST("", CustomerController.SaveDataCustomer)
	// api.PUT("/", CustomerController.EditDataCustomer)
	api.POST("/list", CustomerController.ListDataCustomerByName)
	// api.POST("/check/supplier", MerchantController.CheckOrderMerchantSupplier)

	api = r.Group("/api/supplier")
	api.POST("/page/:page/count/:count", SupplierController.FilterDataSupplier)
	api.POST("", SupplierController.SaveDataSupplier)
	api.PUT("", SupplierController.EditDataSupplier)

	BrandController := new(controllers.BrandController)
	brand := r.Group("/api/brand")
	brand.POST("/page/:page/count/:count", BrandController.GetBrand)
	brand.GET("/id/:id", BrandController.GetFilterBrand)
	brand.POST("", BrandController.SaveBrand)
	brand.PUT("", BrandController.UpdateBrand)
	brand.GET("", BrandController.GetBrandLike)

	ProductController := new(controllers.ProductController)
	product := r.Group("/api/product")
	product.POST("/page/:page/count/:count", ProductController.GetProductListPaging)
	product.GET("/id/:id", ProductController.GetProductDetails)
	product.POST("", ProductController.SaveProduct)
	product.GET("/list", ProductController.ProductList)
	product.GET("", ProductController.GetProductLike)

	ProductGroupController := new(controllers.ProductGroupController)
	productGroup := r.Group("/api/product-group")
	productGroup.POST("/page/:page/count/:count", ProductGroupController.GetProductGroupPaging)
	productGroup.GET("/id/:id", ProductGroupController.GetProductGroupDetails)
	productGroup.POST("", ProductGroupController.SaveProductGroup)
	productGroup.PUT("", ProductGroupController.UpdateProductGroup)

	LookupController := new(controllers.LookupController)
	lookup := r.Group("/api/lookup")
	lookup.GET("", LookupController.GetLookupByGroup)
	lookup.POST("/page/:page/count/:count", LookupController.GetLookupPaging)
	lookup.GET("/id/:id", LookupController.GetLookupFilter)
	lookup.GET("/name/:name", LookupController.GetLookupGroupName)
	lookup.GET("/group", LookupController.GetDistinctLookup)
	lookup.POST("", LookupController.SaveLookup)
	// lookup.PUT("/", LookupController.UpdateLookup)

	LookupGroupController := new(controllers.LookupGroupController)
	lookupGroup := r.Group("/api/lookup-group")
	lookupGroup.GET("", LookupGroupController.GetLookupGroup)

	RoleController := new(controllers.RoleController)
	api = r.Group("/api/role")
	api.POST("/filter/page/:page/count/:count", RoleController.GetRole)
	api.POST("/", RoleController.SaveRole)
	api.PUT("/", RoleController.UpdateRole)

	AccMatrixController := new(controllers.AccessMatrixController)
	MenuController := new(controllers.MenuController)
	api = r.Group("/api/menu")
	api.GET("/list-user-menu", cekToken, MenuController.GetMenuByUser)
	api.GET("/list-all-active-menu", AccMatrixController.GetAllActiveMenu)
	api.GET("/role/:roleId", AccMatrixController.GetMenuByRoleID)
	api.POST("/role/:roleId", AccMatrixController.SaveRoleMenu)

	api = r.Group("/api/sales-order")
	api.GET("/:id", OrderController.GetByOrderId)
	api.POST("/page/:page/count/:count", cekToken, OrderController.FilterData)
	api.POST("", cekToken, OrderController.Save)
	api.POST("/approve", cekToken, OrderController.Approve)
	api.POST("/reject", cekToken, OrderController.Reject)
	api.POST("/invoice/:id", OrderController.PrintInvoice)

	api = r.Group("/api/sales-order-detail")
	api.POST("/page/:page/count/:count", OrderDetailController.GetDetail)
	api.POST("", cekToken, OrderDetailController.Save)
	api.DELETE("/:id", OrderDetailController.DeleteById)

	// Dashboard
	dashboard := r.Group("/dashboard")
	dashboard.POST("/order-qty", DashboardController.FilterDataDashboard)

	return r

}

func cekSignature(c *gin.Context) {

	fmt.Println("cek signature")
	timestamp := c.Request.Header.Get("timestamp")
	signature := c.Request.Header.Get("signature")

	// 5H5GTtcehHqOLDgIzNu8
	key := beego.AppConfig.DefaultString("secret.key", "")
	// body := c.Request.Body
	body := "{}"
	res := dto.LoginResponseDto{}

	if ret := security.ValidateSignature(timestamp, key, signature, body); ret != true {
		res.ErrCode = kons.ERR_CODE_54
		res.ErrDesc = kons.ERR_CODE_54_MSG
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
	}
}

func cekToken(c *gin.Context) {

	res := models.Response{}
	tokenString := c.Request.Header.Get("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") == false {
		res.ErrCode = kons.ERR_CODE_54
		res.ErrDesc = kons.ERR_CODE_54_MSG
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			res.ErrCode = kons.ERR_CODE_54
			res.ErrDesc = kons.ERR_CODE_54_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			// return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(kons.TokenSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("claims : ", claims)

		fmt.Println("User name from TOKEN ", claims["user"])

		unixNano := time.Now().UnixNano()
		timeNowInInt := unixNano / 1000000

		tokenCreated := (claims["tokenCreated"])
		dto.CurrUser = (claims["user"]).(string)
		currUserId := (claims["userId"]).(string)
		dto.CurrUserId, _ = strconv.ParseInt(currUserId, 10, 64)

		fmt.Println("now : ", timeNowInInt)
		fmt.Println("token created time : ", tokenCreated)
		fmt.Println("user by token : ", dto.CurrUser)
		fmt.Println("user by token ID : ", dto.CurrUserId)

		tokenCreatedInString := tokenCreated.(string)
		tokenCreatedInInt, errTokenExpired := strconv.ParseInt(tokenCreatedInString, 10, 64)

		if errTokenExpired != nil {
			res.ErrCode = kons.ERR_CODE_55
			res.ErrDesc = kons.ERR_CODE_55_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if ((timeNowInInt - tokenCreatedInInt) / 1000) > kons.TokenExpiredInMinutes {
			res.ErrCode = kons.ERR_CODE_55
			res.ErrDesc = kons.ERR_CODE_55_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		fmt.Println("Token already used for ", (timeNowInInt-tokenCreatedInInt)/1000, "sec, Max expired ", kons.TokenExpiredInMinutes, "sec ")
		// fmt.Println("token Valid ")

	} else {
		res.ErrCode = kons.ERR_CODE_54
		res.ErrDesc = kons.ERR_CODE_54_MSG
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
}

// CORSMiddleware ...
// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.Method == "OPTIONS" {
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
// 			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
// 			c.Writer.Header().Set("Content-Type", "application/json, charset=utf-8")
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// // @APIVersion 1.0.0
// // @Title beego Test API
// // @Description beego has a very cool tools to autogenerate documents for your API
// // @Contact astaxie@gmail.com
// // @TermsOfServiceUrl http://beego.me/
// // @License Apache 2.0
// // @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// package routers

// import (
// 	"distribution-system-be/controllers"

// 	"github.com/astaxie/beego"
// )

// func init() {
// 	ns := beego.NewNamespace("/v1",

// 		beego.NSNamespace("/object",
// 			beego.NSInclude(
// 				&controllers.ObjectController{},
// 			),
// 		),
// 		beego.NSNamespace("/user",
// 			beego.NSInclude(
// 				&controllers.UserController{},
// 			),
// 		),
// 	)
// 	beego.AddNamespace(ns)
// }
