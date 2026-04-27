package user

import "database/sql"

type UserRepo interface {

}

type UserRepos struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) UserRepo {
	return &UserRepos{
		db: db,
	}
}