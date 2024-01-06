package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://DraconDev:s8JrG9azEWXN@ep-summer-credit-29657807.us-east-2.aws.neon.tech/Main?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)
}
