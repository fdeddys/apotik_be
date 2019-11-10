package routers

import (
	"fmt"
	"net/http"
	"oasis-be/controllers"
	"oasis-be/models"
	dto "oasis-be/models/dto"
	"oasis-be/utils/security"
	"strconv"
	"strings"
	"time"

	kons "oasis-be/constants"

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
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	UserController := new(controllers.UserController)
	MerchantController := new(controllers.MerchantController)
	SupplierController := new(controllers.SupplierController)
	IssuerController := new(controllers.IssuerController)
	OrderController := new(controllers.OrderController)
	VendorController := new(controllers.VendorController)
	OrderDetailController := new(controllers.OrderDetailController)
	SupplierGroupController := new(controllers.SupplierGroupController)
	OrderStatusController := new(controllers.OrderStatusController)
	DashboardController := new(controllers.DashboardController)
	HistoryVendorController := new(controllers.HistoryVendorController)
	FollowUpOrderController := new(controllers.FollowUpOrderController)
	NooDocController := new(controllers.NooDocController)

	api := r.Group("/user")
	api.POST("/page/:page/count/:count", cekToken, UserController.GetUser)
	api.POST("/", cekToken, UserController.SaveDataUser)
	api.PUT("/", cekToken, UserController.UpdateUser)

	AuthController := new(controllers.AuthController)
	api = r.Group("/auth")
	api.POST("/login", AuthController.Login)
	api.GET("/get_cur_user", cekToken, AuthController.GetCurrPass)
	api.POST("/change_pass", cekToken, AuthController.ChangePass)

	api = r.Group("/merchant")
	api.POST("/page/:page/count/:count", cekToken, MerchantController.FilterDataMerchant)
	api.POST("/", cekToken, MerchantController.SaveDataMerchant)
	api.PUT("/", cekToken, MerchantController.EditDataMerchant)
	api.POST("/list", cekToken, MerchantController.ListDataMerchant)
	api.GET("", cekToken, MerchantController.ListDataMerchantByName)
	// api.POST("/check/supplier", MerchantController.CheckOrderMerchantSupplier)

	api = r.Group("/supplier")
	api.POST("/page/:page/count/:count", cekToken, SupplierController.FilterDataSupplier)
	api.POST("/", cekToken, SupplierController.SaveDataSupplier)
	api.PUT("/", cekToken, SupplierController.EditDataSupplier)
	api.POST("/detail/:supplier_id", cekToken, SupplierController.ListDataSupplierById)
	api.POST("/upload", cekToken, SupplierController.UploadImageSupplier)

	// route supplier merchant
	api.POST("/page/:page/count/:count/merchant/:supplier_id", cekToken, SupplierController.FilterDataSupplierMerchant)
	api.PUT("/merchant/:supplier_id", cekToken, SupplierController.UpdateDataSupplierMerchant)
	api.POST("/merchant/:supplier_id", cekToken, SupplierController.SaveDataSupplierMerchant)

	// route supplier warehouse
	api.POST("/page/:page/count/:count/warehouse/:supplier_id", cekToken, SupplierController.FilterDataSupplierWarehouse)
	api.POST("/warehouse/:supplier_id", cekToken, SupplierController.SaveDataSupplierWarehouse)
	api.PUT("/warehouse/:supplier_id", cekToken, SupplierController.UpdateDataSupplierWarehouse)
	api.POST("/warehouse", cekToken, SupplierController.GetDataWarehouseBySupplierId)

	// route supplier product
	api.POST("/page/:page/count/:count/product/:supplier_id", cekToken, SupplierController.FilterDataSupplierPrice)
	api.POST("/product/:supplier_id", cekToken, SupplierController.SaveDataSupplierPrice)
	api.PUT("/product/:supplier_id", cekToken, SupplierController.UpdateDataSupplierPrice)

	// route supplier noo checklist
	api.POST("/page/:page/count/:count/noo_checklist/:supplier_id", cekToken, SupplierController.FilterDataSupplierNooChecklist)
	api.POST("/noo_checklist/:supplier_id", cekToken, SupplierController.SaveDataSupplierNooChecklist)
	api.PUT("/noo_checklist/:supplier_id", cekToken, SupplierController.UpdateDataSupplierNooChecklist)

	// route supplier noo doc
	api.POST("/page/:page/count/:count/approve/noo_doc/:supplier_id", cekToken, SupplierController.FilterDataSupplierNooDocApprove)
	// api.POST("/page/:page/count/:count/noo_doc/:supplier_id", cekToken, SupplierController.FilterDataSupplierNooDoc)
	// api.POST("/noo_doc/:supplier_id", cekToken, SupplierController.SaveDataSupplierNooDoc)
	// api.PUT("/noo_doc/:supplier_id", cekToken, SupplierController.UpdateDataSupplierNooDoc)

	// route upload picture
	api.POST("/ktp/upload", cekToken, SupplierController.UploadImageKtp)
	api.POST("/npwp/upload", cekToken, SupplierController.UploadImageNpwp)
	api.POST("/merchant_picture", cekToken, SupplierController.SaveDataMerchantPicture)
	api.POST("/merchant_picture/:merchant_code", cekToken, SupplierController.ListMerchantPicture)

	// route api for sfa
	api.POST("/list", SupplierController.ListDataSupplier)
	api.POST("/list/products", SupplierController.GetDataListProductBySupplierId)
	api.POST("/check/merchant", SupplierController.CheckMerchantBySupplier)
	api.POST("/register/merchant", SupplierController.SubmitNOO)
	api.POST("/approve/NOO", SupplierController.ApproveNOO)

	api = r.Group("/issuer")
	api.POST("/page/:page/count/:count", cekToken, IssuerController.FilterDataIssuer)
	api.POST("/", cekToken, IssuerController.SaveDataIssuer)
	api.PUT("/", cekToken, IssuerController.EditDataIssuer)
	api.POST("/list", cekToken, IssuerController.ListDataIssuer)
	api.GET("", cekToken, IssuerController.ListDataIssuerByName)

	BrandController := new(controllers.BrandController)
	brand := r.Group("/brand")
	brand.POST("/page/:page/count/:count", cekToken, BrandController.GetBrand)
	brand.GET("/id/:id", cekToken, BrandController.GetFilterBrand)
	brand.POST("/", cekToken, BrandController.SaveBrand)
	brand.PUT("/", cekToken, BrandController.UpdateBrand)
	brand.GET("", cekToken, BrandController.GetBrandLike)

	ProductController := new(controllers.ProductController)
	product := r.Group("/product")
	product.POST("/page/:page/count/:count", cekToken, ProductController.GetProductListPaging)
	product.GET("/id/:id", cekToken, ProductController.GetProductDetails)
	product.POST("/", cekToken, ProductController.SaveProduct)
	product.PUT("/", cekToken, ProductController.UpdateProduct)
	product.GET("/list", cekToken, ProductController.ProductList)
	product.POST("/upload", cekToken, ProductController.UploadImage)
	product.GET("", cekToken, ProductController.GetProductLike)

	ProductGroupController := new(controllers.ProductGroupController)
	productGroup := r.Group("/productgroup")
	productGroup.POST("/page/:page/count/:count", cekToken, ProductGroupController.GetProductGroupPaging)
	productGroup.GET("/id/:id", cekToken, ProductGroupController.GetProductGroupDetails)
	productGroup.POST("/", cekToken, ProductGroupController.SaveProductGroup)
	productGroup.PUT("/", cekToken, ProductGroupController.UpdateProductGroup)

	LookupController := new(controllers.LookupController)
	lookupGroup := r.Group("/lookup")
	lookupGroup.GET("", cekToken, LookupController.GetLookupByGroup)
	lookupGroup.POST("/page/:page/count/:count", cekToken, LookupController.GetLookupPaging)
	lookupGroup.GET("/id/:id", cekToken, LookupController.GetLookupFilter)
	lookupGroup.GET("/group", cekToken, LookupController.GetDistinctLookup)
	lookupGroup.POST("/", cekToken, LookupController.SaveLookup)
	lookupGroup.PUT("/", cekToken, LookupController.UpdateLookup)

	RoleController := new(controllers.RoleController)
	api = r.Group("/role")
	api.POST("/page/:page/count/:count", cekToken, RoleController.GetRole)
	api.POST("/", cekToken, RoleController.SaveRole)
	api.PUT("/", cekToken, RoleController.UpdateRole)

	AccMatrixController := new(controllers.AccessMatrixController)
	MenuController := new(controllers.MenuController)
	api = r.Group("/menu")
	api.GET("/list-user-menu", cekToken, MenuController.GetMenuByUser)
	api.GET("/list-all-active-menu", cekToken, AccMatrixController.GetAllActiveMenu)
	api.GET("/role/:roleId", cekToken, AccMatrixController.GetMenuByRoleID)
	api.POST("/role/:roleId", cekToken, AccMatrixController.SaveRoleMenu)

	api = r.Group("/order")
	api.POST("/page/:page/count/:count", cekToken, OrderController.FilterData)
	api.POST("/save-so", cekToken, OrderController.SaveSO)
	api.POST("/get-detail/page/:page/count/:count", cekToken, OrderDetailController.GetDetail)
	api.POST("/autodebet", cekToken, OrderController.Autodebet)
	api.POST("/release-so", cekToken, OrderController.ReleaseSO)
	api.POST("/manual-pay", cekToken, OrderController.ManualPay)
	api.POST("/get-list-status/page/:page/count/:count", cekToken, OrderStatusController.GetListStatus)
	api.POST("/invoice", cekToken, OrderController.PrintInvoice)
	api.POST("/reject-so", cekToken, OrderController.RejectSO)
	api.POST("/follow-up/page/:page/count/:count", cekToken, FollowUpOrderController.GetFollowUpOrder)

	api = r.Group("/api/vendor")
	api.POST("/login", VendorController.Login)
	api.GET("/getso", VendorController.GetSo)
	api.POST("/update-status", VendorController.UpdateStatus)

	// SupplierGroup
	supplierGroup := r.Group("/suppliergroup")
	supplierGroup.POST("/page/:page/count/:count", cekToken, SupplierGroupController.GetSupplierGroupPaging)
	supplierGroup.GET("/id/:id", cekToken, SupplierGroupController.GetSupplierGroupDetails)
	supplierGroup.POST("/", cekToken, SupplierGroupController.SaveSupplierGroup)
	supplierGroup.PUT("/", cekToken, SupplierGroupController.UpdateSupplierGroup)
	supplierGroup.GET("/", cekToken, SupplierGroupController.GetListSupplierGroup)

	// Dashboard
	dashboard := r.Group("/dashboard")
	dashboard.POST("/order-qty", cekToken, DashboardController.FilterDataDashboard)

	//History Vendor
	historyVendor := r.Group("/history")
	historyVendor.POST("page/:page/count/:count", cekToken, HistoryVendorController.GetHistoryVendorPaging)

	// Subscriber
	apiSubscriber := r.Group("/subscriber")
	apiSubscriber.POST("/order/page/:page/count/:count", cekSignature, OrderController.FilterData)

	api = r.Group("/noo-doc")
	api.POST("/page/:page/count/:count", cekToken, NooDocController.FilterDataNooDoc)
	api.POST("/approve", cekToken, NooDocController.GetStatusApproveNooByMerchantAndSupplier)
	api.PUT("/", cekToken, NooDocController.UpdateDataNooDoc)

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

		fmt.Println("now : ", timeNowInInt)
		fmt.Println("token created time : ", tokenCreated)
		fmt.Println("user by token : ", dto.CurrUser)

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
// 	"oasis-be/controllers"

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
