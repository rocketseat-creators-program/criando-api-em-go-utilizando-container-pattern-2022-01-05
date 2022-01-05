package services

import (
	"github.com/GuilhermeCaruso/expertscrud/src/interfaces"
	"github.com/GuilhermeCaruso/expertscrud/src/repositories"
	"github.com/GuilhermeCaruso/expertscrud/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
