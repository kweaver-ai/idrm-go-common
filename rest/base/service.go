package base

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const schema = `http://`

var Service *ServiceConfig

func init() {
	Service = &ServiceConfig{}
	Service.Init()
}

func (s *ServiceConfig) SetDebug(remoteHost string) {
	sv := reflect.ValueOf(s)
	sv = sv.Elem()
	for i := 0; i < sv.NumField(); i++ {
		sv.Field(i).SetString(remoteHost)
	}
}

// Init 初始化服务注册，省去了诸多服务名的注册，更新chart的麻烦
func (s *ServiceConfig) Init() {
	st := reflect.TypeOf(s)
	st = st.Elem()
	sv := reflect.ValueOf(s)
	sv = sv.Elem()

	for i := 0; i < st.NumField(); i++ {
		fieldType := st.Field(i)
		fieldEnvTag := fieldType.Tag.Get("env")
		if fieldEnvTag == "" {
			panic("missing service env tag")
		}
		//解析host tag 赋值服务地址
		fieldHostTag := fieldType.Tag.Get("host")
		if fieldHostTag != "" {
			fmt.Printf("key %v using struct tag host: %v\n", fieldEnvTag, fieldHostTag)
			addr := fixSchema(fieldHostTag)
			sv.Field(i).SetString(addr)
		}
		//设置环境变量
		existEnvValue := os.Getenv(fieldEnvTag)
		if existEnvValue != "" {
			existEnvValue = fixSchema(existEnvValue)
			fmt.Printf("key %v using exists env: %v\n", fieldEnvTag, existEnvValue)
			sv.Field(i).SetString(existEnvValue)
		}

	}
}

func (s *ServiceConfig) ACHosts() [][]string {
	st := reflect.TypeOf(s).Elem()
	sv := reflect.ValueOf(s).Elem()

	res := make([][]string, 0)
	for i := 0; i < st.NumField(); i++ {
		fieldType := st.Field(i)
		//判断该服务是不是有权限
		permissionTag := fieldType.Tag.Get("permission")
		if permissionTag == "" || permissionTag != "true" {
			continue
		}
		//解析host tag 赋值服务地址
		fieldHostTag := fieldType.Tag.Get("host")
		if fieldHostTag == "" {
			continue
		}
		fieldValue := sv.Field(i).String()
		if fieldValue == "" {
			continue
		}
		//到这里就是符合条件的服务了
		fvs := strings.Split(fieldHostTag, ":")
		if len(fvs) != 2 {
			continue
		}
		res = append(res, []string{
			fvs[0],
			fixSchema(fieldValue),
		})
	}
	return res
}

func (s *ServiceConfig) GetServiceHostByName(name string) string {
	st := reflect.TypeOf(s)
	st = st.Elem()

	for i := 0; i < st.NumField(); i++ {
		fieldType := st.Field(i)
		//解析host tag 赋值服务地址
		fieldHostTag := fieldType.Tag.Get("host")
		if fieldHostTag == "" {
			continue
		}
		if strings.Contains(fieldHostTag, name+":") {
			return fixSchema(fieldHostTag)
		}
	}
	return ""
}

// ServiceConfig 服务地址注册
type ServiceConfig struct {
	AuthServiceHost            string `env:"AUTH_SERVICE_HOST" host:"auth-service:8155"`
	UserMgnHost                string `env:"USER_MANAGE_HOST" host:"user-management-private:30980"`
	UserMgnPublicHost          string `env:"USER_MANAGE_PUBLIC_HOST" host:"user-management-public:30980"`
	HydraAdminHost             string `env:"HYDRA_HOST" host:"hydra-admin:4445"`
	HydraPublicHost            string `env:"HYDRA_PUBLIC_HOST" host:"hydra-public:4444"`
	ConfigurationCenterHost    string `env:"CONFIGURATION_CENTER_HOST" host:"configuration-center:8133"`
	BusinessGroomingHost       string `env:"BUSINESS_GROOMING_HOST" host:"business-grooming:8123"  permission:"false"`
	IndicatorManagement        string `env:"INDICATOR_MANAGEMENT_HOST" host:"indicator-management:8213"`
	TaskCenterHost             string `env:"TASK_CENTER_HOST" host:"task-center:8143"  permission:"false"`
	DataCatalogHost            string `env:"DATA_CATALOG_HOST" host:"data-catalog:8153"  permission:"false"`
	DataSubjectHost            string `env:"DATA_SUBJECT_HOST" host:"data-subject:8123"  permission:"false"`
	DataViewHost               string `env:"DATA_VIEW_HOST" host:"data-view:8123"  permission:"false"`
	MetaDataHost               string `env:"METADATA_HOST" host:"af-vega-metadata:80"`
	DataApplicationServiceHost string `env:"DATA_APPLICATION_SERVICE_HOST" host:"data-application-service:8156"`
	StandardizationHost        string `env:"STANDARDIZATION_HOST" host:"standardization:80"`
	WorkFlowRestHost           string `env:"WORKFLOW_REST_HOST" host:"workflow-rest:9800"`
	DocAuditRestHost           string `env:"DOC_AUDIT_REST_HOST" host:"doc-audit-rest:9800"`
	// AnyRobot host, such as https://anyrobot.example.org
	AnyRobotHost         string `env:"ANYROBOT_HOST"`
	VirtualEngineHost    string `env:"VIRTUAL_ENGINE_HOST" host:"af-vega-gateway:8099"`
	AfSailorHost         string `env:"AF_SAILOR_HOST" host:"af-sailor:9797"`
	BasicBigDataHost     string `env:"BASIC_BIGDATA_HOST" host:"basic-bigdata-service:8287"`
	DataSyncHost         string `env:"DATA_SYNC_CONNECT_HOST" host:"data-sync-connect:80"`
	DemandManagementHost string `env:"DEMAND_MANAGEMENT_HOST" host:"demand-management:8280"   permission:"false"`
	// basic-search
	BasicSearchHost          string `env:"BASIC_SEARCH_HOST" host:"basic-search:8163"`
	SceneAnalysisHost        string `env:"SCENE_ANALYSIS_HOST" host:"scene-analysis:8193"`
	AfSailorServiceHost      string `env:"AF_SAILOR_SERVICE_HOST" host:"af-sailor-service:80"`
	AuthorizationPublicHost  string `env:"AUTHORIZATION_PUBLIC_HOST" host:"authorization-public:30920"`
	AuthorizationPrivateHost string `env:"AUTHORIZATION_PRIVATE_HOST" host:"authorization-private:30920"`
}

func fixSchema(s string) string {
	if strings.HasPrefix(s, schema) {
		return s
	}
	return schema + s
}

func RemoveSchema(s string) string {
	if strings.HasPrefix(s, schema) {
		return strings.TrimPrefix(s, schema)
	}
	return s
}
