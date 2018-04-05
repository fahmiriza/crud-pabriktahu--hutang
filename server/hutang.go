package server

import (
	"context"
)

type hutang struct {
	writer ReadWriter
}

func NewHutang(writer ReadWriter) HutangService {
	return &hutang{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *hutang) AddHutangService(ctx context.Context, hutang Hutang) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddHutang(hutang)
	if err != nil {
		return err
	}

	return nil
}

func (c *hutang) ReadHutangService(ctx context.Context) (Hutangs, error) {
	kar, err := c.writer.ReadHutang()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *hutang) UpdateHutangService(ctx context.Context, kar Hutang) error {
	err := c.writer.UpdateHutang(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *hutang) ReadHutangByNamaService(ctx context.Context, namahutang string) (Hutang, error) {
	kar, err := c.writer.ReadHutangByNama(namahutang)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
