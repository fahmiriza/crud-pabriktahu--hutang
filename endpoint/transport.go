package endpoint

import (
	"context"

	scv "MiniProject/git.bluebird.id/mini/hutang/server"

	pb "MiniProject/git.bluebird.id/mini/hutang/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcHutangServer struct {
	addHutang        grpctransport.Handler
	readHutangByNama grpctransport.Handler
	readHutang       grpctransport.Handler
	updateHutang     grpctransport.Handler
}

func NewGRPCHutangServer(endpoints HutangEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.HutangServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcHutangServer{
		addHutang: grpctransport.NewServer(endpoints.AddHutangEndpoint,
			decodeAddHutangRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddHutang", logger)))...),
		readHutangByNama: grpctransport.NewServer(endpoints.ReadHutangByNamaEndpoint,
			decodeReadHutangByNamaRequest,
			encodeReadHutangByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadHutangByNama", logger)))...),
		readHutang: grpctransport.NewServer(endpoints.ReadHutangEndpoint,
			decodeReadHutangRequest,
			encodeReadHutangResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadHutang", logger)))...),
		updateHutang: grpctransport.NewServer(endpoints.UpdateHutangEndpoint,
			decodeUpdateHutangRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateHutang", logger)))...),
	}
}

func decodeAddHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddHutangReq)
	return scv.Hutang{KodeHutang: req.GetKodeHutang(), NamaHutang: req.GetNamaHutang(), TanggalHutang: req.GetTanggalHutang(), TotalHutang: req.GetTotalHutang(), Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcHutangServer) AddHutang(ctx oldcontext.Context, hutang *pb.AddHutangReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addHutang.ServeGRPC(ctx, hutang)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadHutangByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadHutangByNamaReq)
	return scv.Hutang{NamaHutang: req.NamaHutang}, nil
}

func decodeReadHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadHutangByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Hutang)
	return &pb.ReadHutangByNamaResp{KodeHutang: resp.KodeHutang, NamaHutang: resp.NamaHutang, TanggalHutang: resp.TanggalHutang, TotalHutang: resp.TotalHutang, Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func encodeReadHutangResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Hutangs)

	rsp := &pb.ReadHutangResp{}

	for _, v := range resp {
		itm := &pb.ReadHutangByNamaResp{
			KodeHutang:    v.KodeHutang,
			NamaHutang:    v.NamaHutang,
			TanggalHutang: v.TanggalHutang,
			TotalHutang:   v.TotalHutang,
			Status:        v.Status,
			CreateBy:      v.CreateBy,
		}
		rsp.AllHutang = append(rsp.AllHutang, itm)
	}
	return rsp, nil
}

func (s *grpcHutangServer) ReadHutangByNama(ctx oldcontext.Context, namahutang *pb.ReadHutangByNamaReq) (*pb.ReadHutangByNamaResp, error) {
	_, resp, err := s.readHutangByNama.ServeGRPC(ctx, namahutang)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadHutangByNamaResp), nil
}

func (s *grpcHutangServer) ReadHutang(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadHutangResp, error) {
	_, resp, err := s.readHutang.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadHutangResp), nil
}

func decodeUpdateHutangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateHutangReq)
	return scv.Hutang{KodeHutang: req.KodeHutang, NamaHutang: req.NamaHutang, TanggalHutang: req.TanggalHutang, TotalHutang: req.TotalHutang, Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func (s *grpcHutangServer) UpdateHutang(ctx oldcontext.Context, cus *pb.UpdateHutangReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateHutang.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
