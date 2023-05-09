package db

import (
	"context"
	"os"

	"memcache/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func NewPostgreSQL() (*PostgreSQL, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}

func (p *PostgreSQL) FindByNConst(nconst string) (models.Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res models.Name

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return models.Name{}, err
	}

	return res, nil
}
