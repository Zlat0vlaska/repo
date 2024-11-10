package handlers

import (
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/contracts/dto"
	"reliab-test/internal/contracts/utils"

	"github.com/gin-gonic/gin"
)

type VacancyHandler struct {
	log            *slog.Logger
	vacancyService vacancyServiceInf
}

func BuildVacancyHandler(vacancyService vacancyServiceInf, log *slog.Logger) *VacancyHandler {
	return &VacancyHandler{vacancyService: vacancyService, log: log}
}

func (vh *VacancyHandler) GetVacancies(c *gin.Context) {
	ctx := c.Request.Context()
	var requestDto dto.GetVacanciesRequest

	err := c.ShouldBind(&requestDto)
	if err != nil {
		vh.log.Error("Wrong data sent", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	vh.log.Info("Data got:", requestDto)

	vacancies, err := vh.vacancyService.GetVacancies(ctx, dto.BuildGetVacanciesDtoToFilter(requestDto))

	if err != nil {
		vh.log.Error("Failed to get all vacancies", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	vacanciesDto := dto.BuildVacancyToGetVacancyResponseMultiple(vacancies)

	vh.log.Info("Data sent", vacanciesDto)
	utils.BuildSuccessResponse(c, vacanciesDto)
}

func (vh *VacancyHandler) GetVacancyByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	vh.log.Info("Data got:", id)

	if id == "" {
		vh.log.Error("Missing ID in request")
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	vacancy, err := vh.vacancyService.GetVacancyByID(ctx, id)
	if err != nil {
		vh.log.Error("Failed to get vacancy by ID", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	if vacancy == nil {
		utils.BuildErrorResponse(c, http.StatusNotFound, domain_errors.VacancyNotFound)
		return
	}

	response := dto.BuildVacancyToGetVacancyResponse(*vacancy)

	utils.BuildSuccessResponse(c, response)
}
