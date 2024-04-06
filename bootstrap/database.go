package bootstrap

import (
	"context"
	"fmt"

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
func CreateTables(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS cars (
			reg_num TEXT PRIMARY KEY,
			mark TEXT NOT NULL,
			model TEXT NOT NULL,
			year INTEGER NOT NULL,
			owner_name TEXT NOT NULL,
			owner_surname TEXT NOT NULL,
			owner_patronymic TEXT NOT NULL
		);
	`)
	return err
}
