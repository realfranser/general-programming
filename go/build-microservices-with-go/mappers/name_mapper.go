package mappers

import "microservice_example/models"

func NameModelToGorm(model models.Name) *models.GormNames {
	return &models.GormNames{
		NConst:    model.NConst,
		Name:      model.Name,
		BirthYear: model.BirthYear,
		DeathYear: model.DeathYear,
	}
}

func NameGormToModel(gormName models.GormNames) *models.Name {
	return &models.Name{
		NConst:    gormName.NConst,
		Name:      gormName.Name,
		BirthYear: gormName.BirthYear,
		DeathYear: gormName.DeathYear,
	}
}
