package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func conectarBanco() {
	stringdeconexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(stringdeconexao))
	if err != nil {
		panic("Falha ao conectar com o banco de dados")
	}
}
