package dao

import (
	"github.com/jinzhu/gorm"
	xerrors "github.com/pkg/errors"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var dbErr error
	db, dbErr = gorm.Open("")
	if dbErr != nil {
		err = xerrors.Wrap(dbErr, "dao: new db failed")
		return
	}
	cf = func() { db.Close() }
	return
}
