package configuration_center

import "github.com/kweaver-ai/idrm-go-frame/core/enum"

type SourceType enum.Object

const NotClassified = "not_classified"

var (
	InfoSystem    = enum.New[SourceType](1, "info_system")
	DataWarehouse = enum.New[SourceType](2, "data_warehouse")
)

func SourceTypeStringToInt(sourceType string) int32 {
	return enum.ToInteger[SourceType](sourceType).Int32()
}

type UserStatus int32

const (
	UserNormal UserStatus = 1
	UserDelete UserStatus = 2
)

// sszd-open  省市直达开关枚举
const (
	SSZDOpenTrue  = "true"
	SSZDOpenFalse = "false"
)

const (
	DataDictQueryTypeAll    = ""  //查询所有的字典
	DataDictQueryTypeSSZD   = "1" //查询省市直达
	DataDictQueryTypeNormal = "0" //查询产品字典
)
