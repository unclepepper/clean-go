package postgres

import (
	"github.com/unclepepper/clean-go/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgDb struct {
	*gorm.DB
}

func NewPostgres(conf *configs.Config) *PgDb {
	db, err := gorm.Open(postgres.Open(conf.Db.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &PgDb{db}
}
