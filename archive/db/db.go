package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func InitDB() error {
	var err error

	databaseUrl := "postgres://postgres:postgres@localhost:5432/records" //os.Getenv("DATABASE_URL")
	db, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}

	return nil // wie war das ncoh, wann gibt man errors usw zurück? Oder ist das eh ne unsitte -> nachgucken error handling und exception handling
}

func GetDB() *pgx.Conn {
	return db
}
