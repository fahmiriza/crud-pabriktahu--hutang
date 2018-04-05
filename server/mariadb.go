package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addHutang          = `insert into hutang(kodehutang, namahutang,tanggalhutang,totalhutang, status, createby,createon)values (?,?,?,?,?,?,?)`
	selectHutang       = `select kodehutang, namahutang,tanggalhutang,totalhutang, createby from hutang Where status =1`
	updateHutang       = `update hutang set namahutang=?,tanggalhutang=?,totalhutang=?, status=?, updateby=?, updateon=? where kodehutang=?`
	selectHutangByNama = `select kodehutang, namahutang,tanggalhutang,totalhutang, status, createon from hutang where namahutang=?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddHutang(hutang Hutang) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addHutang, hutang.KodeHutang, hutang.NamaHutang, hutang.TanggalHutang, hutang.TotalHutang, OnAdd, hutang.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadHutangByNama(namahutang string) (Hutang, error) {
	fmt.Println("show by nama")
	hutang := Hutang{NamaHutang: namahutang}
	err := rw.db.QueryRow(selectHutangByNama, namahutang).Scan(&hutang.KodeHutang, &hutang.NamaHutang, &hutang.TanggalHutang, &hutang.TotalHutang, &hutang.Status, &hutang.CreateBy)

	if err != nil {
		return Hutang{}, err
	}

	return hutang, nil
}

func (rw *dbReadWriter) ReadHutang() (Hutangs, error) {
	fmt.Println("show all")
	hutang := Hutangs{}
	rows, _ := rw.db.Query(selectHutang)
	defer rows.Close()
	for rows.Next() {
		var k Hutang
		err := rows.Scan(&k.KodeHutang, &k.NamaHutang, &k.TanggalHutang, &k.TotalHutang, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return hutang, err
		}
		hutang = append(hutang, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return hutang, nil
}

func (rw *dbReadWriter) UpdateHutang(kar Hutang) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateHutang, kar.NamaHutang, kar.TanggalHutang, kar.TotalHutang, kar.Status, kar.UpdateBy, time.Now(), kar.KodeHutang)

	fmt.Println("name:", kar.NamaHutang, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
