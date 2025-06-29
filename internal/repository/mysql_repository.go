package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLRepository struct {
	DB *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{
		DB: db,
	}
}

func InitDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
