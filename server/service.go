package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Hutang.PabrikTahu.id"
	OnAdd     Status = 1
)

type Hutang struct {
	KodeHutang    string
	NamaHutang    string
	TanggalHutang string
	TotalHutang   string
	Status        int32
	CreateBy      string
	UpdateBy      string
}
type Hutangs []Hutang

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddHutang(Hutang) error
	ReadHutang() (Hutangs, error)
	UpdateHutang(Hutang) error
	ReadHutangByNama(string) (Hutang, error)
}

type HutangService interface {
	AddHutangService(context.Context, Hutang) error
	ReadHutangService(context.Context) (Hutangs, error)
	UpdateHutangService(context.Context, Hutang) error
	ReadHutangByNamaService(context.Context, string) (Hutang, error)
}
