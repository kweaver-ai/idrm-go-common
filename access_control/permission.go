package access_control

// 追加额外的访问权限控制
func AddExtraAccessControl(userRoles []string, entry *ScopeTransfer) (res *AccessControlRes) {
	res = new(AccessControlRes)
	res.InfoResourceCatalogFrontEnd = 1
	for _, role := range userRoles {
		if role != TCSystemMgm {
			res.InfoResourceCatalogAudit = 1
		}
		if role == TCDataOperationEngineer || role == TCDataDevelopmentEngineer {
			res.InfoResourceCatalogCataloging = 1
		}
	}
	res.ScopeTransfer = *entry
	return
}

type AccessControlRes struct {
	ScopeTransfer
	InfoResourceCatalogFrontEnd   uint8 `json:"info_resource_catalog_frontend"`   // 信息资源目录前台（数据服务超市）
	InfoResourceCatalogCataloging uint8 `json:"info_resource_catalog_cataloging"` // 信息资源目录编目
	InfoResourceCatalogAudit      uint8 `json:"info_resource_catalog_audit"`      // 信息资源目录审核 // [/]
}
