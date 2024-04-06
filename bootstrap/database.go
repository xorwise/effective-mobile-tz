package bootstrap

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewDatabaseConnection(env *Env) (*pgx.Conn, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", env.PostgresUser, env.PostgresPassword, env.PostgresHost, env.PostgresPort, env.PostgresDatabase)
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CloseDatabaseConnection(conn *pgx.Conn) {
	conn.Close(context.Background())
}

// Function for creating tables at application start
func MigrateDatabase(conn *pgx.Conn) error {
	path := "migrations/1_create_cars_table.up.sql"
	c, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sql := string(c)
	_, err = conn.Exec(context.Background(), sql)
	return err
}
