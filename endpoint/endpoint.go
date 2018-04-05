package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/mini/hutang/server"

	kit "github.com/go-kit/kit/endpoint"
)

type HutangEndpoint struct {
	AddHutangEndpoint        kit.Endpoint
	ReadHutangByNamaEndpoint kit.Endpoint
	ReadHutangEndpoint       kit.Endpoint
	UpdateHutangEndpoint     kit.Endpoint
}

func NewHutangEndpoint(service svc.HutangService) HutangEndpoint {
	addHutangEp := makeAddHutangEndpoint(service)
	readHutangByNamaEp := makeReadHutangByNamaEndpoint(service)
	readHutangEp := makeReadHutangEndpoint(service)
	updateHutangEp := makeUpdateHutangEndpoint(service)
	return HutangEndpoint{AddHutangEndpoint: addHutangEp,
		ReadHutangByNamaEndpoint: readHutangByNamaEp,
		ReadHutangEndpoint:       readHutangEp,
		UpdateHutangEndpoint:     updateHutangEp,
	}
}

func makeAddHutangEndpoint(service svc.HutangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Hutang)
		err := service.AddHutangService(ctx, req)
		return nil, err
	}
}

func makeReadHutangByNamaEndpoint(service svc.HutangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Hutang)
		result, err := service.ReadHutangByNamaService(ctx, req.NamaHutang)
		return result, err
	}
}

func makeReadHutangEndpoint(service svc.HutangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadHutangService(ctx)
		return result, err
	}
}

func makeUpdateHutangEndpoint(service svc.HutangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Hutang)
		err := service.UpdateHutangService(ctx, req)
		return nil, err
	}
}
