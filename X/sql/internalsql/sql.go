package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweet/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		cfg.Db_user,
		cfg.Db_pass,
		cfg.Db_host,
		cfg.Db_port,
		cfg.Db_name,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("database connected")

	return db, nil
}