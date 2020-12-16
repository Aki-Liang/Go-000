package dao

import (
	"context"

	"homework04/internal/model"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	xerrors "github.com/pkg/errors"
)

var Provider = wire.NewSet(New, NewDB)

type Dao interface {
	Close()
	Ping(ctx context.Context) error
	Find(ctx context.Context, target string) (*model.Target, error)
}

type dao struct {
	db *gorm.DB
}

func New(db *gorm.DB) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *gorm.DB) (d Dao, cf func(), err error) {
	d = &dao{
		db: db,
	}
	cf = d.Close
	return
}

func (d *dao) Close() {
	//do nothing
}

func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

func (d *dao) Find(ctx context.Context, target string) (*model.Target, error) {
	var t model.Target
	err := d.db.Table("target_table").Where("target_name=?", target).First(&t).Error
	if err != nil {
		return nil, xerrors.Wrap(err, "dao: find target failed")
	}
	return &t, nil
}
