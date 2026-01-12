package business_grooming

type TableNameCheckResp struct {
	Name   string `json:"name"`
	Repeat bool   `json:"repeat"`
}
type DataTableFieldInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	RawType     string `json:"rawType"`
	OrigType    string `json:"origType"`
}

type FormPathInfoResp struct {
	FormID             string `json:"form_id"`              //业务表ID
	FormName           string `json:"form_name"`            //业务表名称
	BusinessModelID    string `json:"business_model_id"`    //业务模型ID
	BusinessModelName  string `json:"business_model_name"`  //模型名称
	MainBusinessID     string `json:"main_business_id"`     //主干业务ID，以前的业务流程
	MainBusinessName   string `json:"main_business_name"`   //主干业务名称，以前的业务流程
	DomainID           string `json:"domain_id" `           //业务域id
	DomainName         string `json:"domain_name"`          //业务域名称
	DomainGroupID      string `json:"domain_group_id"`      //业务域分组id
	DomainGroupName    string `json:"domain_group_name"`    //业务域分组名称
	ProcessPathID      string `json:"process_path_id"`      //包含主干业务ID的path
	ProcessPathName    string `json:"process_path_name"`    //包含主干业务名称的path
	DepartmentID       string `json:"department_id"`        //业务表部门
	DepartmentName     string `json:"department_name"`      //部门名称
	DepartmentPathID   string `json:"department_path_id"`   //业务表部门
	DepartmentPathName string `json:"department_path_name"` //部门名称
}

// MainBusinessModelResp 主干业务及其关联的业务模型信息（扁平结构）
type MainBusinessModelResp struct {
	MainBusinessID   string `json:"main_business_id"`   //主干业务ID
	MainBusinessName string `json:"main_business_name"` //主干业务名称
	ModelID          string `json:"model_id"`           //业务模型ID
	ModelName        string `json:"model_name"`         //业务模型名称
}
