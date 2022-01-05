package repositories

import (
	"github.com/GuilhermeCaruso/expertscrud/src/interfaces"
	"github.com/GuilhermeCaruso/expertscrud/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
