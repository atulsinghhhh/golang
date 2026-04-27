package internalsql

import (
	"database/sql"
	"go-tweet/internal/config"
)

func ConnectMySQL(cfg *config.Config)(*sql.DB)