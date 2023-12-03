package databases

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func InitDatabase(dsn string) (err error) {
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
