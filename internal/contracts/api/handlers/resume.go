package handlers

import (
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/contracts/dto"
	"reliab-test/internal/contracts/utils"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct { 
	log           *slog.Logger
	resumeService resumeService
}

func BuildResumeHandler(resumeService resumeService, log *slog.Logger) *ResumeHandler {
	return &ResumeHandler{resumeService: resumeService, log: log}
}

func (rh *ResumeHandler) GetAllResumes(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.GetAllResumes

	if err := c.ShouldBindQuery(&params); err != nil {
		rh.log.Error("Failed to bind query parameters", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}
	rh.log.Info("Data got:", params)

	filter := dto.GetResumesDtoToFilter(params)
	resumes, totalResumes, totalPages, err := rh.resumeService.GetAllResumes(ctx, filter)
	if err != nil {
		rh.log.Error("Failed to get all resumes", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}
	if totalResumes == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No resumes found"})
		return
	}
	//totalPages := (totalResumes + limit - 1) / limit
	resumesDto := dto.BuildResumeToGetResponseMultiple(resumes)

	response := gin.H{
		"results":      resumesDto,
		"totalPages":   totalPages,
		"totalResults": totalResumes,
	}

	rh.log.Info("Data sent", response)
	utils.BuildSuccessResponse(c, response)
}

func (rh *ResumeHandler) GetResumeByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	rh.log.Info("Data got:", id)

	if id == "" {
		rh.log.Error("Missing ID in request")
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	resumeId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		rh.log.Error("Invalid ID format", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	resume, err := rh.resumeService.GetResumeByID(ctx, resumeId)
	if err != nil {
		rh.log.Error("Failed to get resume by ID", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	if resume == nil {
		utils.BuildErrorResponse(c, http.StatusNotFound, domain_errors.ResumeNotFound)
		return
	}

	dto := dto.BuildResumeToGetResponse(*resume)

	utils.BuildSuccessResponse(c, dto)
}
