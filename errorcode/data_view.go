package errorcode

func init() {
	RegisterErrorCode(dataViewErrorMap)
}

const dataViewPreCoder = "ServiceCall.DataView."
const (
	CascadeDeleteSubjectDomainViewRelatedError = dataViewPreCoder + "CascadeDeleteSubjectDomainViewRelatedError"
	QueryViewCountError                        = dataViewPreCoder + "QueryViewCountError"
	GetDataViewDetailsError                    = dataViewPreCoder + "GetDataViewDetailsError"
	GetDataViewFieldError                      = dataViewPreCoder + "GetDataViewFieldError"
)

var dataViewErrorMap = ErrorCode{
	CascadeDeleteSubjectDomainViewRelatedError: {
		Description: "刪除主题域下绑定的视图关联失败",
		Solution:    "",
	},
	QueryViewCountError: {
		Description: "查询视图数量错误",
		Solution:    "请联系管理员",
	},
	GetDataViewDetailsError: {
		Description: "查询视图详情错误",
		Solution:    "请联系管理员",
	},
	GetDataViewFieldError: {
		Description: "查询视图字段错误",
		Solution:    "请联系管理员",
	},
}
