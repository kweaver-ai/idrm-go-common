package gocephclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	cephClientOnce      sync.Once
	cephClientSingleton CephClient
)

type CephClient interface {
	Upload(id string, data []byte) (err error)
	Down(id string) (data []byte, err error)
	Preview(id, fileName string) (info *UploadInfo, err error)
	Delete(id string) (err error)
}

type FormField struct{}

// type Headers struct {
// 	Authorization string `json:"Authorization"`
// 	ContentType   string `json:"Content-Type"`
// 	XAmzDate      string `json:"x-amz-date"`
// 	XObsDate      string `json:"x-obs-date"`
// }

type Headers map[string]string

type UploadInfo struct {
	FormField FormField `json:"form_field"`
	Headers   Headers   `json:"headers"`
	Method    string    `json:"method"`
	Url       string    `json:"url"`
}

type cephClient struct {
	addr       string // oss addr
	protocol   string // oss protocol
	is_default string // oss_is_default
	bucket     string // oss bucket
	app        string // oss app
	httpClient *http.Client
}

type bucketInfo struct {
	AccessId           string `json:"accessId"`
	AccessKey          string `json:"accessKey"`
	CdnName            string `json:"cdnName"`
	HttpPort           int    `json:"httpPort"`
	HttpsPort          int    `json:"httpsPort"`
	InternalServerName string `json:"internalServerName"`
	Name               string `json:"name"`
	Provider           string `json:"provider"`
	ProviderDetail     string `json:"providerDetail"`
	ServerName         string `json:"serverName"`
}
type StorageInfo struct {
	StorageName    string     `json:"storageName"`
	StorageId      string     `json:"storageId"`
	SiteId         string     `json:"siteId"`
	OssgwPort      string     `json:"ossgwPort"`
	OssgwHttpsPort string     `json:"ossgwHttpsPort"`
	OssgwHost      string     `json:"ossgwHost"`
	IsDefault      bool       `json:"isDefault"`
	IsCacheBucket  bool       `json:"isCacheBucket"`
	InternalgwPort string     `json:"internalgwPort"`
	InternalgwHost string     `json:"internalgwHost"`
	HasOSSGW       bool       `json:"hasOSSGW"`
	Enabled        bool       `json:"enabled"`
	App            string     `json:"app"`
	T              bucketInfo `json:"bucketInfo"`
}

func getOssHost() string {
	content, err := ioutil.ReadFile("./config.txt")
	if err == nil {
		string_slice := strings.Split(string(content), ":")
		return string_slice[1] + ":" + string_slice[2]
	}
	return ""
}

func NewCephClient(httpClient *http.Client) (CephClient, error) {
	var err error
	cephClientOnce.Do(func() {
		var app, addr, protocol, is_default, bucket string
		if app = os.Getenv("OSS_APP"); app == "" {
			err = errors.New("need to set $OSS_APP")
			return
		}
		if addr = os.Getenv("OSS_HOST"); addr == "" {
			err = errors.New("need to set $OSS_HOST")
			return
		}
		if protocol = os.Getenv("OSS_PROTOCOL"); protocol == "" {
			err = errors.New("need to set $OSS_PROTOCOL")
			return
		}

		if is_default = os.Getenv("OSS_IS_DEFAULT"); is_default == "" {
			err = errors.New("need to set $IS_DEFAULT")
			return
		}
		if bucket = os.Getenv("OSS_BUCKET"); bucket == "" {
			err = errors.New("need to set $BUCKET")
			return
		}
		cephClientSingleton = &cephClient{
			app:        app,
			addr:       addr,
			protocol:   protocol,
			is_default: is_default,
			bucket:     bucket,
			httpClient: httpClient,
		}
	})

	return cephClientSingleton, err
}

func (c *cephClient) SignRequest(req *http.Request, info *UploadInfo) {
	for k, v := range info.Headers {
		req.Header.Set(k, v)
	}
	// req.Header.Set("x-amz-date", info.Headers.XAmzDate)
	// req.Header.Set("Authorization", info.Headers.Authorization)
	// req.Header.Set("Content-Type", info.Headers.ContentType)
}

func (c *cephClient) CreateRequest(objID string, method string, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	return req
}

func GetStorageId(addr string, isDefault string, app string, bucket string, protocol string) string {
	url := protocol + "://" + addr + "/api/ossgateway/v1/objectstorageinfo?isCache=false"
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(reqest)
	if err != nil {
		fmt.Printf("faild to get storage info %v err %v", url, err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := make([]StorageInfo, 0)
	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Printf("fail to unmarshal body %v, err: %v", string(body), err.Error())
		return ""
	}
	for i := range res {
		is_defalut := strconv.FormatBool(res[i].IsDefault)
		if is_defalut == isDefault && res[i].T.Name == bucket || res[i].App == app {
			return res[i].StorageId
		}
	}
	return ""
}

func GetInfo(storage_id string, file string, method string, addr string, fileName string) (*UploadInfo, error) {
	var url string
	if method == "upload" {
		url = "http" + "://" + addr + "/api/ossgateway/v1/upload/" + storage_id + "/" + file + "?request_method=PUT"
	} else if method == "download" {
		url = "http" + "://" + addr + "/api/ossgateway/v1/download/" + storage_id + "/" + file
	} else if method == "preview" {
		expireTime := fmt.Sprintf("%d", time.Now().Add(24*time.Hour).Unix())
		url = "http" + "://" + addr + "/api/ossgateway/v1/download/" + storage_id + "/" + file + "?type=query_string&Expires=" + expireTime + "&response-content-type=application/pdf"
	} else if method == "delete" {
		url = "http" + "://" + addr + "/api/ossgateway/v1/delete/" + storage_id + "/" + file
	} else {
		return nil, errors.New("get info only support get or head")
	}
	fmt.Printf("storage url %v", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 预览时，safari直接播放mp3音频和mp4视频，需要存储返回的content-type明确为audio/mp3和video/mp4；此处需要判断
	if method == "preview" {
		if strings.ToLower(GetExtendName(fileName)) == ".mp4" {
			req.Header.Add("response-content-type", "video/mp4")
		} else if strings.ToLower(GetExtendName(fileName)) == ".mp3" {
			req.Header.Add("response-content-type", "audio/mp3")
		} else {
			//req.Header.Add("response-content-type", "application/pdf")
			//req.Header.Add("response-content-disposition", "inline")
			req.Header.Add("Content-Type", "application/json;charset=utf-8")
		}
	} else {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ossgateway response %v", string(body))
	res := new(UploadInfo)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

/*
1、get storage_id from oss;
2、get upload info from oss
3、put byte to oss
*/
func (c *cephClient) Upload(id string, data []byte) (err error) {
	if data == nil {
		var errorEmptyData = errors.New("data cannot be null")
		fmt.Printf("id %v, data cannot be null %v", id, string(data))
		return errorEmptyData
	}
	addr := c.addr
	isDefault := c.is_default
	app := c.app
	bucket := c.bucket
	protocol := c.protocol
	if addr == "" {
		var errorConfig = errors.New("fail get addr from config")
		fmt.Printf("id %v, err: %v", id, errorConfig)
		return errorConfig
	}
	storage_id := GetStorageId(addr, isDefault, app, bucket, protocol)
	if storage_id == "" {
		fmt.Println("step1: addr: %v, fail get storage id from oss", addr)
		var errorStorayId = errors.New("fail get storage id from oss, check  configuration")
		return errorStorayId
	}
	res_info, err := GetInfo(storage_id, id, "upload", addr, "")
	if err != nil {
		fmt.Printf("fail get put info from oss,error:%v", err.Error())
		var errorGetInfo = errors.New("fail get put info from oss")
		return errorGetInfo
	}
	req := c.CreateRequest(id, res_info.Method, res_info.Url)
	req.ContentLength = int64(len(data))
	body := bytes.NewBuffer(data)
	req.Body = ioutil.NopCloser(body)
	c.SignRequest(req, res_info)
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("fail to do request error:%v", err.Error())
		return err
	}
	defer res.Body.Close()

	respContent, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Upload respnse body %v", string(respContent))

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("ERROR: GetDocConfig: Response = %v", res.StatusCode)
		return
	}
	return
}

func (c *cephClient) Down(id string) (content []byte, err error) {
	addr := c.addr
	if addr == "" {
		var errorConfig = errors.New("fail get addr from config")
		return nil, errorConfig
	}
	isDefault := c.is_default
	app := c.app
	bucket := c.bucket
	protocol := c.protocol
	storage_id := GetStorageId(addr, isDefault, app, bucket, protocol)
	if storage_id == "" {
		var errorStorayId = errors.New("fail get storage id from oss")
		return nil, errorStorayId
	}
	res_info, err := GetInfo(storage_id, id, "download", addr, "")
	if err != nil {
		fmt.Printf("fail get download info from oss error:%v", err.Error())
		var errorGetInfo = errors.New("fail get download info from oss")
		return nil, errorGetInfo
	}

	req := c.CreateRequest(id, res_info.Method, res_info.Url)
	c.SignRequest(req, res_info)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("failed to  do request error:%v", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("ERROR: GetDocConfig: Response = %v", resp.StatusCode)
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *cephClient) Preview(id, saveName string) (info *UploadInfo, err error) {
	addr := c.addr
	if addr == "" {
		var errorConfig = errors.New("fail get addr from config")
		return nil, errorConfig
	}
	isDefault := c.is_default
	app := c.app
	bucket := c.bucket
	protocol := c.protocol
	storage_id := GetStorageId(addr, isDefault, app, bucket, protocol)
	if storage_id == "" {
		var errorStorayId = errors.New("fail get storage id from oss")
		return nil, errorStorayId
	}
	res_info, err := GetInfo(storage_id, id, "preview", addr, saveName)
	if err != nil {
		fmt.Printf("fail get preview info from oss error:%v", err.Error())
		var errorGetInfo = errors.New("fail get preview info from oss")
		return nil, errorGetInfo
	}

	return res_info, nil
}

func GetExtendName(filePath string) (extName string) {
	filenameWithSuffix := path.Base(path.Clean(filePath))
	extName = path.Ext(filenameWithSuffix)
	if extName != ".gz" {
		return
	}
	ext := path.Ext(strings.TrimSuffix(filenameWithSuffix, extName))
	if ext == ".tar" {
		extName = ".tar.gz"
	}
	return
}

func (c *cephClient) Delete(id string) (err error) {
	addr := c.addr
	if addr == "" {
		var errorConfig = errors.New("fail get addr from config")
		return errorConfig
	}
	isDefault := c.is_default
	app := c.app
	bucket := c.bucket
	protocol := c.protocol
	storage_id := GetStorageId(addr, isDefault, app, bucket, protocol)
	if storage_id == "" {
		var errorStorayId = errors.New("fail get storage id from oss")
		return errorStorayId
	}
	res_info, err := GetInfo(storage_id, id, "delete", addr, "")
	if err != nil {
		fmt.Printf("fail get delete info from oss error:%v", err.Error())
		var errorGetInfo = errors.New("fail get delete info from oss")
		return errorGetInfo
	}

	req := c.CreateRequest(id, res_info.Method, res_info.Url)
	c.SignRequest(req, res_info)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("failed to  do request error:%v", err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("ERROR: GetDocConfig: Response = %v", resp.StatusCode)
		return err
	}
	return
}
