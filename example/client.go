package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.bluebird.id/mini/hutang/endpoint"
	opt "MiniProject/git.bluebird.id/mini/util/grpc"
	util "MiniProject/git.bluebird.id/mini/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCHutangClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Hutang
	//client.AddHutangService(context.Background(), svc.Hutang{KodeHutang: "079", NamaHutang: "damia", TanggalHutang: "21 maret 2014", TotalHutang: "5000", CreateBy: "Admin"})

	//Get Hutang By Nama
	//cusNama, _ := client.ReadHutangByNamaService(context.Background(), "Riza")
	//fmt.Println("hutang based on namahutang:", cusNama)

	//List Hutang
	cuss, _ := client.ReadHutangService(context.Background())
	fmt.Println("all hutangs:", cuss)

	//Update Karyawan
	//client.UpdateHutangService(context.Background(), svc.Hutang{KodeHutang: "079", NamaHutang: "Riza", TanggalHutang: "21 maret 2014", TotalHutang: "5000", CreateBy: "Admin"})

}
