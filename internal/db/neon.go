package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// func Neon() {
// 	connStr := "postgresql://DraconDev:s8JrG9azEWXN@ep-summer-credit-29657807.us-east-2.aws.neon.tech/Main?sslmode=require"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	var version string
// 	if err := db.QueryRow("select version()").Scan(&version); err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("version=%s\n", version)
// }

type Database struct {
	*sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping to db: %v", err)
	}
	return &Database{db}, nil
}

func (db *Database) GetVersion() (string, error) {
	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		return "", fmt.Errorf("could not get version: %v", err)
	}
	return version, nil
}
