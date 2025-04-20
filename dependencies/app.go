package dependencies

import (
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"gorm.io/gorm"
)

type Dependency struct {
	Repository *repositories.Repository
}

func NewDependency(db *gorm.DB) *Dependency {

	repo := &repositories.Repository{
		DB: db,
	}

	return &Dependency{
		Repository: repo,
	}
}