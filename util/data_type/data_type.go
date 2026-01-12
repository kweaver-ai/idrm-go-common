package constant

//region 业务大类型

const SimpleChar = "char"
const SimpleInt = "int"
const SimpleFloat = "float"
const SimpleDecimal = "decimal"
const SimpleBool = "bool"
const SimpleDate = "date"
const SimpleDatetime = "datetime"
const SimpleTime = "time"
const SimpleBinary = "binary"
const SimpleOther = "other"

//endregion

//region 业务大类型 转 字符枚举值

var SimpleType2ChMapping = map[string]string{
	SimpleChar:     "字符型",
	SimpleInt:      "整数型",
	SimpleFloat:    "小数型",
	SimpleDecimal:  "高精度型",
	SimpleBool:     "布尔型",
	SimpleDate:     "日期型",
	SimpleDatetime: "日期时间型",
	SimpleTime:     "时间型",
	SimpleBinary:   "二进制型",
	SimpleOther:    "未定义型",
}

//endregion

//region 字符枚举值 转 业务大类型

var Ch2SimpleTypeMapping = map[string]string{
	"字符型":   SimpleChar,
	"字符串":   SimpleChar,
	"整数型":   SimpleInt,
	"小数型":   SimpleFloat,
	"高精度型":  SimpleDecimal,
	"布尔型":   SimpleBool,
	"日期型":   SimpleDate,
	"日期":    SimpleDate,
	"日期时间型": SimpleDatetime,
	"时间型":   SimpleTime,
	"二进制型":  SimpleBinary,
	"未定义型":  SimpleOther,
	"图片":    SimpleOther,
}

//endregion
