package main

import (
	"context"
	"github.com/jdw-art/tiktok_commence/demo/demo_proto/biz/service"
	pbapi "github.com/jdw-art/tiktok_commence/demo/demo_proto/kitex_gen/pbapi"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
