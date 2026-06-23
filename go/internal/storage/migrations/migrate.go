package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedFS embed.FS

// Run ejecuta todas las migraciones SQL embebidas
func Run(db *sql.DB) error {
	goose.SetBaseFS(embedFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("fallo seteando dialecto postgres: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("error al ejecutar migraciones: %w", err)
	}

	return nil
}
