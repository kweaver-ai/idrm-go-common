package basic_bigdata_service

import (
	"context"

	"github.com/kweaver-ai/idrm-go-frame/core/enum"
)

type Driven interface {
	// 上传文件
	UploadFile(ctx context.Context, req *UploadReq) (res *UploadResp, err error)
	// 删除文件
	DeleteFile(ctx context.Context, req *DeleteReq) (err error)
	// 根据ossId 删除文件
	DeleteOssFile(ctx context.Context, req *DeleteReq) (err error)
	// 根据关联对象id批量删除文件
	DeleteFiles(ctx context.Context, req *DeleteReq) (err error)
	// 分页查询文件列表
	QueryPageFile(ctx context.Context, req *QueryPageReq) (res *QueryResp, err error)
	// 下载文件
	DownloadFile(ctx context.Context, req *DownloadReq) (res *DownloadResp, err error)
	// 根据关联对象id获取对应文件数量
	GetCountsByIds(ctx context.Context, req *GetCountsByIdsReq) (res *GetCountsByIdsResp, err error)
	// 根据关联对象id批量删除不包括这些文件ID的文件
	DeleteExcludeFiles(ctx context.Context, req *DeleteExcludeOssIdsReq) (err error)
	// 预览文件链接
	PreviewPdfHref(ctx context.Context, req *PreviewPdfReq) (res *PreviewPdfResp, err error)

	GetLabelByIds(ctx context.Context, ids []string) (res *GetLabelByIdsRes, err error)
}

type UploadReq struct {
	Name            string `json:"name"  binding:"required" example:"1.docx"`                                                                                           // 文件名称(包括后缀名)
	Type            string `json:"type"  binding:"required,oneof=standard_spec 3_def_response construct_basis construct_content file_resource" example:"standard_spec"` // 文件类型standard_spec标准规范，3_def_response三定职责，construct_basis建设依据，construct_content建设内容，file_resource文件资源
	RelatedObjectID string `json:"related_object_id" binding:"required"  example:"1"`                                                                                   // 关联对象ID数组（部门ID、业务领域id、信息系统id）
	OssID           string `json:"oss_id"  binding:"required"  example:"1"`                                                                                             // 此项不为空则用它下载文件
	Content         []byte `json:"content"`                                                                                                                             // OssID为空时此项生效(字节流方式)
	FileSize        int64  `json:"file_size"  binding:"omitempty"  example:"1"`                                                                                         //文件大小
}

type UploadResp struct {
	OssID string `json:"oss_id"`
}

type DownloadReq struct {
	ID    string `form:"id"`
	OssID string `form:"oss_id"`
}

type DeleteExcludeOssIdsReq struct {
	RelatedObjectId string `json:"related_object_id"  form:"related_object_id"  binding:"required"  example:"12222"`   //关联对象Id
	ExcludeOssIDs   string `json:"exclude_oss_ids"  form:"exclude_oss_ids"  binding:"required"  example:"12222,33333"` //排除的ossID,多个用“,”分割
}

type DeleteReq struct {
	ID string `uri:"id"`
}

type QueryPageReq struct {
	Keyword          string   `json:"keyword" form:"keyword" binding:"VerifyXssString,omitempty,min=1,max=128"`                                                                        // 关键字查询，字符无限制
	Type             string   `json:"type"  form:"type" binding:"required,oneof=standard_spec 3_def_response construct_basis construct_content file_resource" example:"standard_spec"` // 文件类型standard_spec标准规范，3_def_response三定职责，construct_basis建设依据，construct_content建设内容，file_resource文件资源
	IDs              []string `json:"id"  form:"id" binding:"omitempty,lte=20,dive,max=36"`                                                                                            //文件ID数组，最多20个，文件雪花id
	RelatedObjectIDs []string `json:"related_object_id"  form:"related_object_id" binding:"omitempty,lte=30,dive,max=36"`                                                              // 关联对象ID数组（部门ID、业务领域id、信息系统id），最多30个
	Offset           int      `json:"offset" form:"offset,default=1" binding:"omitempty,min=1" default:"1"  example:"1"`                                                               // 页码，默认1
	Limit            int      `json:"limit" form:"limit,default=12" binding:"omitempty,min=1,max=120" default:"12"  example:"12"`                                                      // 每页大小，默认12
}

type FileInfo struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Size            int64  `json:"size"`
	RelatedObjectID string `json:"related_object_id"`
	ExportOssID     string `json:"export_oss_id"`
	PreviewOssID    string `json:"preview_oss_id"`
	CreatedAt       int64  `json:"created_at"`
}

type QueryResp struct {
	TotalCount int         `json:"total_count"`
	Entries    []*FileInfo `json:"entries"`
}

type DownloadResp struct {
	Name    string
	Content []byte
}

type GetCountsByIdsReq struct {
	Type string   `json:"type"  binding:"required,oneof=standard_spec 3_def_response construct_basis construct_content file_resource" example:"standard_spec"` // 文件类型standard_spec标准规范，3_def_response三定职责，construct_basis建设依据，construct_content建设内容，file_resource文件资源
	IDs  []string `json:"ids" binding:"required"`                                                                                                              // ID列表
}

type GetCountsByIdsResp struct {
	Entries []*CountInfo `json:"entries"`
}

type CountInfo struct {
	ID    string `json:"id" binding:"required"`    // 关联对象id
	Count int64  `json:"count" binding:"required"` // 关联文件的数量
}

type PreviewPdfReq struct {
	ID        string `form:"id"  binding:"required"  example:"12"`            //文件id
	PreviewID string `form:"preview_id"  binding:"required"  example:"12222"` //文件预览对象存储id
}

type PreviewPdfResp struct {
	ID        string `json:"id"  binding:"required"  example:"12"`                                      //文件id
	PreviewID string `json:"preview_id"  binding:"required"  example:"12222"`                           //文件预览对象存储id
	HrefUrl   string `json:"href_url"  binding:"required"  example:"http://xxx.xxx.xxx/ddd/sss?ss.pdf"` //预览链接地址
}

// [办公文档文件类型]
type EnumOfficeDocumentFileType enum.Object

var (
	EnumStandardSpecification      = enum.New[EnumOfficeDocumentFileType](1, "standard_spec", "标准规范")
	EnumThreeDefinedResponsibility = enum.New[EnumOfficeDocumentFileType](2, "3_def_response", "三定职责")
	EnumConstructionBasis          = enum.New[EnumOfficeDocumentFileType](3, "construct_basis", "建设依据")
	EnumConstructionContent        = enum.New[EnumOfficeDocumentFileType](4, "construct_content", "建设内容")
	EnumFileResource               = enum.New[EnumOfficeDocumentFileType](5, "file_resource", "文件资源")
) // [/]

type GetLabelByIdsRes struct {
	// LabelResp 标签列表
	LabelResp []LabelItem `json:"label_resp"`
}

// LabelItem 标签项
type LabelItem struct {
	ID string `json:":"id"`
	// Name 标签名称
	Name string `json:"name"`
	// Path 标签路径
	Path string `json:"path"`
}
