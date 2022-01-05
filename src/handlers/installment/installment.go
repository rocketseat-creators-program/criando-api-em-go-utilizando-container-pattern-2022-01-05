package installment

import (
	"net/http"

	"github.com/GuilhermeCaruso/expertscrud/src/interfaces"
	"github.com/GuilhermeCaruso/expertscrud/src/services"
	"github.com/GuilhermeCaruso/expertscrud/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, serviceContainer services.ServiceContainer) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		installmentService: serviceContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	group := ih.router.Group("/installment")

	group.Post("", ih.CreateInstallment)
}

func (ih InstallmentHandler) CreateInstallment(c *fiber.Ctx) error {
	var body structs.Installment

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Reponse{
			Data: err.Error(),
			Tag:  "BAD_REQUEST",
		})
	}

	err = ih.installmentService.Create(body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.Reponse{
			Data: err.Error(),
			Tag:  "INTERNAL_SERVER_ERROR",
		})
	}

	return c.Status(http.StatusCreated).JSON(structs.Reponse{
		Data: "installment criada!",
	})
}
