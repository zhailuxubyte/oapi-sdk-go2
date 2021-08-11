// Code generated by lark suite oapi sdk gen
package v2

import (
	"github.com/larksuite/oapi-sdk-go"
)

type Service struct {
	conf    lark.Config
	Tenants *TenantService
}

func NewService(conf lark.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Tenants = newTenantService(s)
	return s
}

type TenantService struct {
	service *Service
}

func newTenantService(service *Service) *TenantService {
	return &TenantService{
		service: service,
	}
}

type TenantQueryReqCall struct {
	ctx     *lark.Context
	tenants *TenantService
	opts    []lark.APIRequestOpt
}

func (rc *TenantQueryReqCall) Do() (*TenantQueryResult, error) {
	var result = &TenantQueryResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/tenant/v2/tenant/query", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.tenants.service.conf, req)
	return result, err
}

func (tenants *TenantService) Query(ctx *lark.Context, opts ...lark.APIRequestOpt) *TenantQueryReqCall {
	return &TenantQueryReqCall{
		ctx:     ctx,
		tenants: tenants,
		opts:    opts,
	}
}