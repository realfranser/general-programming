package models

type Name struct {
	NConst    string `db:"nconst" json:"nconst"`
	Name      string `db:"primary_name" json:"name"`
	BirthYear string `db:"birth_year" json:"birthYear"`
	DeathYear string `db:"death_year" json:"deathYear"`
}

type GormNames struct {
	NConst    string `gorm:"primaryKey;column:nconst"`
	Name      string `gorm:"column:primary_name"`
	BirthYear string
	DeathYear string
}

func (GormNames) TableName() string {
	return "names"
}
