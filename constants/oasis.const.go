package constants

const (
	TokenSecretKey        = "OasI$_sEcrET_key$"
	TokenExpiredInMinutes = 8 * 60 * 60
)

// ERR code Global
const (
	ERR_CODE_00     = "00"
	ERR_CODE_00_MSG = "SUCCESS.."

	ERR_CODE_03     = "03"
	ERR_CODE_03_MSG = "Error, unmarshall body Request"
)

// ERR code Global
const (
	ERR_CODE_30     = "30"
	ERR_CODE_30_MSG = "Failed save data to DB"
)

const (
	ERR_CODE_50     = "50"
	ERR_CODE_50_MSG = "Invalid username / password"

	ERR_CODE_51     = "51"
	ERR_CODE_51_MSG = "Error connection to database"

	ERR_CODE_53     = "53"
	ERR_CODE_53_MSG = "Failed generate token !"

	ERR_CODE_54     = "54"
	ERR_CODE_54_MSG = "Invalid Authorization !"

	ERR_CODE_55     = "55"
	ERR_CODE_55_MSG = "Token expired !"

	ERR_CODE_61     = "61"
	ERR_CODE_61_MSG = "User not found !"

	ERR_CODE_62     = "62"
	ERR_CODE_62_MSG = "Password not match !"

	ERR_CODE_63     = "63"
	ERR_CODE_63_MSG = "Failed Update password !"
)

//nDeskKey ...
const (
	DesKey = "abcdefghijklmnopqrstuvwxyz012345"
)

// ERROR FROM VENDOR
const (
	ERR_CODE_70     = "70"
	ERR_CODE_70_MSG = "ERROR FROM VENDOR "
)

const (
	ERR_CODE_80     = "80"
	ERR_CODE_80_MSG = "Failed save to database"

	ERR_CODE_81     = "81"
	ERR_CODE_81_MSG = "Failed get data from database"
)

const (
	ERR_CODE_90     = "90"
	ERR_CODE_90_MSG = "Failed get from Body Request"
)

// STATUS Sales Order
// 10 = new order
// 20 = approve
// 30 = reject
// 40 = INVOICE
// 50 = PAID

const (
	ERR_CODE_40     = "40"
	ERR_CODE_40_MSG = "Data not found"
)
