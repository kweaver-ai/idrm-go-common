package access_control

//region 以下cs系统内置角色

const ProjectMgm = "09095468-122a-4c6a-8bd3-0c2e438782de"                //项目管理员
const SystemMgm = "033a0ef3-f98f-4498-a8b8-1d6266f500a3"                 //系统管理员
const BusinessMgm = "2927b1fb-df1b-451c-9519-4ba6aa46a22"                //业务管理员
const BusinessOperationEngineer = "3c86b2ff-97e0-4d8b-a904-c8a01fc444fd" //业务运营工程师
const StandardMgmEngineer = "56ce508f-1e1c-4bf6-ba63-1dcbbf980d10"       //标准管理工程师
const DataQualityEngineer = "84f8e148-5571-48c3-9b78-bdd2eecb382d"       //数据质量工程师
const DataAcquisitionEngineer = "9d42f240-b4ad-45e8-ab8b-8727899d3445"   //数据采集工程师
const DataProcessingEngineer = "c7752d11-7aef-4163-a7ba-27f5cf717b8e"    //数据加工工程师
const IndicatorMgmEngineer = "d173ed4c-d4f2-44bb-8a75-646d13e195a2"      //指标管理工程师
//endregion

//region 以下tc系统内置角色

const TCSystemMgm = "00001f64-209f-4260-91f8-c61c6f820136"               //系统管理员
const TCDataOwner = "00002fb7-1e54-4ce1-bc02-626cb1f85f62"               //数据owner
const TCDataButler = "00003148-fbbf-4879-988d-54af7c98c7ed"              //数据管家
const TCDataOperationEngineer = "00004606-f318-450f-bc53-f0720b27acff"   //数据运营工程师
const TCDataDevelopmentEngineer = "00005871-cedd-4216-bde0-94ced210e898" //数据开发工程师
const TCNormal = "0000663b-46a9-45e4-b6f7-a6bd8c18bd46"                  //普通用户
const ApplicationDeveloper = "00007030-4e75-4c5e-aa56-f1bdf7044791"      //应用开发者
const SecurityMgm = "00008516-45b3-44c9-9188-ca656969e20f"               //安全管理员
const PortalMgm = "ab8b1165-5ab8-4054-aa29-2cbb579900f3"                 //门户管理员
//endregion

// 权限配置ID
const ManagerDataView = "982eaf56-74fb-484a-a390-e205d4c80d95"            //管理逻辑视图
const FXGXDJ = "8e7406af-482f-4e6d-ac9e-37b19c69c717"                     //分析和实施供需对接
const ManagerIndicatorPermission = "68e736d6-6a77-4b64-ad89-ead3d6c22c00" // 管理指标
const ManagerDataFLFJPermission = "167d41c2-4b37-47e1-9c29-d103c4873f4f"  //管理数据分类分级
const XQAnalysisPermission = "5c100c9e-5f93-48fb-92ef-d5a898aa3fe0"       //需求分析与实施
const OutBoundApplyPermission = "0077a70c-37c9-46fd-a805-3a4265fb2880"    //处理数据分析需求成果出库
const QRXQZYPermission = "0077a70c-37c9-46fd-a805-3a4265fb2887"           // 确认需求资源

const ManagerDZZZPermission = "0077a70c-37c9-46fd-a805-3a4265fb2874" //管理电子证照目录
