package orm

import (
	"os"

	"microservice_example/mappers"
	"microservice_example/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type PostgreSQLgorm struct {
	db *gorm.DB
}

var dbConnGorm *PostgreSQLgorm

func NewPostgreSQLgorm() (*PostgreSQLgorm, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConnGorm = &PostgreSQLgorm{
		db: db,
	}

	return dbConnGorm, nil
}

func GormFindByNConst(nconst string) (models.Name, error) {
	var result models.GormNames

	tx := dbConnGorm.db.Where("nconst = ?", nconst).First(&result)
	if tx.Error != nil {
		return models.Name{}, tx.Error
	}

	return *mappers.NameGormToModel(result), nil
}
