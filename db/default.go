package db

import (
	"database/sql"
)

var SqlDB *sql.DB

func InitDb() {
	InitRedis()
}
