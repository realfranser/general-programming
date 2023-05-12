package plain_sql

import (
	"context"
	"database/sql"
	"os"

	"microservice_example/models"

	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSQLsql struct {
	pool *sql.DB
}

var dbConnSQL *PostgreSQLsql

func NewPostgreSQLsql() (*PostgreSQLsql, error) {
	pool, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(); err != nil {
		return nil, err
	}

	dbConnSQL = &PostgreSQLsql{
		pool: pool,
	}

	return dbConnSQL, nil
}

func (p *PostgreSQLsql) Close() {
	p.pool.Close()
}

func SQLFindByNConst(nconst string) (models.Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res models.Name

	if err := dbConnSQL.pool.QueryRowContext(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return models.Name{}, err
	}

	return res, nil
}
