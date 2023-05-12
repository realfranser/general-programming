package plain_sql

import (
	"context"
	"microservice_example/models"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSQLsqlx struct {
	db *sqlx.DB
}

var dbConnSQLX *PostgreSQLsqlx

func NewPostgreSQLsqlx() (*PostgreSQLsqlx, error) {
	db, err := sqlx.ConnectContext(context.Background(), "postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	dbConnSQLX = &PostgreSQLsqlx{
		db: db,
	}

	return dbConnSQLX, nil
}

func (p *PostgreSQLsqlx) Close() {
	p.db.Close()
}

func SQLXFindByNConst(nconst string) (models.Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var result models.Name

	if err := dbConnSQLX.db.QueryRowx(query, nconst).StructScan(&result); err != nil {
		return models.Name{}, err
	}

	return result, nil
}
