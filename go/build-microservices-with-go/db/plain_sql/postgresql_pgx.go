package plain_sql

import (
	"context"
	"microservice_example/models"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQLpgx struct {
	pool *pgxpool.Pool
}

var dbConnPGX *PostgreSQLpgx

func NewPostgreSQLpgx() (*PostgreSQLpgx, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	dbConnPGX = &PostgreSQLpgx{
		pool: pool,
	}

	return dbConnPGX, nil
}

func (p *PostgreSQLpgx) Close() {
	p.pool.Close()
}

func PGXFindByNConst(nconst string) (models.Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res models.Name

	if err := dbConnPGX.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return models.Name{}, err
	}

	return res, nil
}
