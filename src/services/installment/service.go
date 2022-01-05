package installment

import (
	"errors"
	"time"

	"github.com/GuilhermeCaruso/expertscrud/src/interfaces"
	"github.com/GuilhermeCaruso/expertscrud/src/repositories"
	"github.com/GuilhermeCaruso/expertscrud/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repositoryContainer repositories.RepositoryContainer) InstallmentService {
	return InstallmentService{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (is InstallmentService) Create(installment structs.Installment) error {

	if installment.DueDay < time.Now().Day() {
		return errors.New("parcela vencida")
	}

	return is.InstallmentRepository.Create(installment)
}
