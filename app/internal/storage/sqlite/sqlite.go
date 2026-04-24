package sqlite

import( 
	
	"database/sql"
	"github.com/atulsinghhhh/golang/internal/config"
	_"github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(Cfg *config.Config)(*Sqlite,error){
	db,err:=sql.Open("sqlite3",Cfg.StoragePath)
	if err!=nil{
		return nil,err
	}

	_,err=db.Exec(`
		CREATE TABLE IF NOT EXIST APP (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			context TEXT
		)
	`)

	if err!=nil{
		return nil,err
	}

	return &Sqlite{
		Db: db,
	},nil

}	


func Storage