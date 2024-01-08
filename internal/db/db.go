package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
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

// add new card to db
func (db *Database) AddCard(text string, tags []string) error {
	// Serialize tags array to JSON, if not empty
	var tagsJSON []byte
	var err error
	if tags != nil {
		tagsJSON, err = json.Marshal(tags)
		if err != nil {
			return err
		}
	}

	// Prepare the SQL query to insert text and tags
	query := "INSERT INTO kanban (name, tags) VALUES ($1, $2)"
	_, err = db.Exec(query, text, tagsJSON)
	if err != nil {
		return err
	}

	return nil
}

// get all cards
func (db *Database) GetAllCards() ([]Card, error) {
	// Prepare the SQL query to get all cards
	query := "SELECT * FROM kanban"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		var tagsJSON []byte
		if err := rows.Scan(&card.ID, &card.Text, &tagsJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(tagsJSON, &card.Tags); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func AccessDB() *Database {
	godotenv.Load(".env")
	connStr := os.Getenv("NEON_LINK")

	myDB, err := NewDatabase(connStr)
	if err != nil {
		panic(err)
	}
	myDB.GetVersion()
	return myDB
}
