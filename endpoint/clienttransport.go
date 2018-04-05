package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.bluebird.id/mini/hutang/server"

	pb "MiniProject/git.bluebird.id/mini/hutang/grpc"

	util "MiniProject/git.bluebird.id/mini/util/grpc"
	disc "MiniProject/git.bluebird.id/mini/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.HutangService"
)

func NewGRPCHutangClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.HutangService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addHutangEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddHutangEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addHutangEp = retry
	}

	var readHutangByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadHutangByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readHutangByNamaEp = retry
	}

	var readHutangEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadHutangEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readHutangEp = retry
	}

	var updateHutangEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateHutang, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateHutangEp = retry
	}
	return HutangEndpoint{AddHutangEndpoint: addHutangEp, ReadHutangByNamaEndpoint: readHutangByNamaEp,
		ReadHutangEndpoint: readHutangEp, UpdateHutangEndpoint: updateHutangEp}, nil
}

func encodeAddHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Hutang)
	return &pb.AddHutangReq{
		KodeHutang:    req.KodeHutang,
		NamaHutang:    req.NamaHutang,
		TanggalHutang: req.TanggalHutang,
		TotalHutang:   req.TotalHutang,
		Status:        req.Status,
		CreateBy:      req.CreateBy,
	}, nil
}

func encodeReadHutangByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Hutang)
	return &pb.ReadHutangByNamaReq{NamaHutang: req.NamaHutang}, nil
}

func encodeReadHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Hutang)
	return &pb.UpdateHutangReq{
		KodeHutang:    req.KodeHutang,
		NamaHutang:    req.NamaHutang,
		TanggalHutang: req.TanggalHutang,
		TotalHutang:   req.TotalHutang,
		Status:        req.Status,
		UpdateBy:      req.UpdateBy,
	}, nil
}

func decodeHutangResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadHutangByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadHutangByNamaResp)
	return svc.Hutang{
		KodeHutang:    resp.KodeHutang,
		NamaHutang:    resp.NamaHutang,
		TanggalHutang: resp.TanggalHutang,
		TotalHutang:   resp.TotalHutang,
		Status:        resp.Status,
		CreateBy:      resp.CreateBy,
	}, nil
}

func decodeReadHutangResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadHutangResp)
	var rsp svc.Hutangs

	for _, v := range resp.AllHutang {
		itm := svc.Hutang{
			KodeHutang:    v.KodeHutang,
			NamaHutang:    v.NamaHutang,
			TanggalHutang: v.TanggalHutang,
			TotalHutang:   v.TotalHutang,
			Status:        v.Status,
			CreateBy:      v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddHutangEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddHutang",
		encodeAddHutangRequest,
		decodeHutangResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddHutang")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddHutang",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadHutangByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadHutangByNama",
		encodeReadHutangByNamaRequest,
		decodeReadHutangByNamaRespones,
		pb.ReadHutangByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadHutangByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadHutangByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadHutangEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadHutang",
		encodeReadHutangRequest,
		decodeReadHutangResponse,
		pb.ReadHutangResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadHutang")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadHutang",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateHutang(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateHutang",
		encodeUpdateHutangRequest,
		decodeHutangResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateHutang")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateHutang",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
