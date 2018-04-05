package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.bluebird.id/mini/hutang/server"
)

func (ke HutangEndpoint) AddHutangService(ctx context.Context, hutang sv.Hutang) error {
	_, err := ke.AddHutangEndpoint(ctx, hutang)
	return err
}

func (ke HutangEndpoint) ReadHutangByNamaService(ctx context.Context, namahutang string) (sv.Hutang, error) {
	req := sv.Hutang{NamaHutang: namahutang}
	fmt.Println(req)
	resp, err := ke.ReadHutangByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Hutang)
	return cus, err
}

func (ke HutangEndpoint) ReadHutangService(ctx context.Context) (sv.Hutangs, error) {
	resp, err := ke.ReadHutangEndpoint(ctx, nil)
	fmt.Println("ke resp y", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Hutangs), err
}

func (ke HutangEndpoint) UpdateHutangService(ctx context.Context, kar sv.Hutang) error {
	_, err := ke.UpdateHutangEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
