package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/domain"
)

var (
	username = "admin"
	password = "admin"
	host     = "localhost"
	port     = "5432"
	dbName   = "clean-architecture-db"
)

func NewPostgreSQLConnection(ctx context.Context) (*bun.DB, error) {
	// Prepare connection string parameterized
	connectionParams := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		dbName,
	)

	// Connect to database doing a PING
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionParams), pgdriver.WithTimeout(time.Second*30)))

	// Verifique se o banco de dados já existe.
	if err := db.Ping(); err != nil {
		log.Printf("erro ao conectar ao banco de dados: %v", err)
	}

	// set connection settings
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Duration(60) * time.Minute)

	dbBun := bun.NewDB(db, pgdialect.New())

	dbBun.RegisterModel((*domain.Order)(nil))

	if _, err := dbBun.NewCreateTable().IfNotExists().Model((*domain.Order)(nil)).Exec(ctx); err != nil {
		return nil, err
	}

	fmt.Println("Db connected")
	return dbBun, nil
}
