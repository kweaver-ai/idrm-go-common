package data_lineage

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

//抽象类，还有三个方法需要具体的实现

type CommonInfoFetcher struct {
	cc configuration_center.Driven
}

func NewInfoFetcher(c configuration_center.Driven) CommonInfoFetcher {
	return CommonInfoFetcher{
		cc: c,
	}
}
func (f CommonInfoFetcher) GetIndicatorInfo(ctx context.Context, id string) (*LineageIndicator, error) {
	panic("implement me")
}

func (f CommonInfoFetcher) GetTableInfo(ctx context.Context, id string) (*TableInfo, error) {
	panic("implement me")
}

func (f CommonInfoFetcher) GetDepartmentInfo(ctx context.Context, id string) (*DepartmentInfo, error) {
	result, err := f.cc.GetDepartmentPrecision(ctx, []string{id})
	if err != nil {
		log.Errorf("query department info %v error %v", id, err.Error())
		return nil, err
	}
	if len(result.Departments) <= 0 {
		log.Errorf("empty department info %v", id)
		return nil, err
	}
	departmentInfo := new(DepartmentInfo)
	copier.Copy(&departmentInfo, &result.Departments[0])
	return departmentInfo, nil
}

func (f CommonInfoFetcher) GetInfoSystem(ctx context.Context, id string) (*InfoSystemInfo, error) {
	infoSystems, err := f.cc.GetInfoSystemsPrecision(ctx, []string{id}, nil)
	if err != nil {
		log.Errorf("query infosystem info %v error %v", id, err.Error())
		return nil, err
	}
	if len(infoSystems) <= 0 {
		log.Errorf("empty infosystem info %v", id)
		return nil, err
	}
	infoSystem := new(InfoSystemInfo)
	copier.Copy(&infoSystem, &infoSystems[0])
	return infoSystem, nil

}

func (f CommonInfoFetcher) GetUserInfo(ctx context.Context, id string) (*UserInfo, error) {
	panic("implement me")
}

func (f CommonInfoFetcher) GetDatasourceInfo(ctx context.Context, id string) (*DataSource, error) {
	panic("implement me")
}
