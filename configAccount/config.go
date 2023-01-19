package configAccount

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	//Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open connection to db
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL_ACCOUNT"))
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses konek DB")
	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
