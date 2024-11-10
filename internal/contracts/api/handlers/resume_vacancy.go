package handlers

import (
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/contracts/dto"
	"reliab-test/internal/contracts/utils"
	"reliab-test/internal/domain"

	"github.com/gin-gonic/gin"
)

type ResumeVacancyHandler struct {
	log                  *slog.Logger
	resumeVacancyService ResumeVacancyService
}

func BuildResumeVacancyHandler(
	resumeVacancyService ResumeVacancyService,
	log *slog.Logger) *ResumeVacancyHandler {
	return &ResumeVacancyHandler{resumeVacancyService: resumeVacancyService, log: log}
}

// CreateLinkHandler обрабатывает создание связи между резюме и вакансией
func (h *ResumeVacancyHandler) CreateLinkHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var reqDto dto.CreateLinkRequest

	if err := c.ShouldBindJSON(&reqDto); err != nil {
		h.log.Error("Can't parse JSON", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	link := dto.BuildDtoToLink(reqDto)

	id, err := h.resumeVacancyService.CreateLink(ctx, link)
	if err != nil {
		h.log.Error("Can't create link resume vacancy", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	respDto := dto.CreateMeetingResponse{ID: id}

	utils.BuildSuccessResponse(c, respDto)
}

// UpdateLinkStatusHandler обрабатывает обновление статуса связи между резюме и вакансией
func (h *ResumeVacancyHandler) UpdateLinkStatusHandler(c *gin.Context) {
	resumeID := c.Param("resume_id")
	vacancyID := c.Param("vacancy_id")
	if vacancyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vacancy_id is required"})
		return
	}
	// Bind the request to dto.UpdateLinkStatusRequest
	var request dto.UpdateLinkStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	// Map dto.UpdateLinkStatusRequest to domain.ResumeVacancy
	linkStatusUpdate := domain.ResumeVacancy{
		ResumeID:      resumeID,
		VacancyID:     vacancyID,
		Status:        request.Status,
		ResumeStatus:  request.ResumeStatus,
		VacancyStatus: request.VacancyStatus,
		Notes:         request.Notes,
	}
	if err := h.resumeVacancyService.UpdateLinkStatus(c.Request.Context(), resumeID, vacancyID, linkStatusUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update link status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link status updated successfully"})
}

// GetResumesByVacancyHandler обрабатывает получение резюме по вакансии
func (h *ResumeVacancyHandler) GetResumesByVacancyHandler(c *gin.Context) {
	vacancyID := c.Param("vacancy_id")

	resumes, err := h.resumeVacancyService.GetResumesByVacancy(c.Request.Context(), vacancyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve resumes"})
		return
	}

	c.JSON(http.StatusOK, resumes)
}

// GetVacanciesByResumeHandler обрабатывает получение вакансий по резюме
func (h *ResumeVacancyHandler) GetVacanciesByResumeHandler(c *gin.Context) {
	h.log.Info("Request Path:", "path", c.Request.URL.Path)
	resumeID := c.Param("resume_id")
	h.log.Info("Received resume_id:", "resume_id", resumeID)

	if resumeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resume_id is required"})
		return
	}
	vacancies, err := h.resumeVacancyService.GetVacanciesByResume(c.Request.Context(), resumeID)
	if err != nil {
		h.log.Error("Error fetching vacancies by resume", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacancies)
}
