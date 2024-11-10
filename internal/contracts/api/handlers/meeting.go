package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/contracts/dto"
	"reliab-test/internal/contracts/utils"
	"reliab-test/internal/services"
)

type MeetingHandler struct {
	log            *slog.Logger
	meetingService *services.MeetingService
}

func BuildMeetingHandler(meetingService *services.MeetingService, log *slog.Logger) *MeetingHandler {
	return &MeetingHandler{meetingService: meetingService, log: log}
}

func (mh *MeetingHandler) CreateMeeting(c *gin.Context) {
	ctx := c.Request.Context()

	var reqDto dto.CreateMeetingRequest

	if err := c.ShouldBindJSON(&reqDto); err != nil {
		mh.log.Error("Can't parse JSON", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	meeting, err := dto.BuildDtoToMeeting(reqDto.MeetingDTO)
	if err != nil {
		mh.log.Error("Can't convert DTO to domain", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	id, err := mh.meetingService.CreateMeeting(ctx, meeting)
	if err != nil {
		mh.log.Error("Can't create meeting", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	respDto := dto.CreateMeetingResponse{ID: id}

	utils.BuildSuccessResponse(c, respDto)
}

func (mh *MeetingHandler) GetMeetings(c *gin.Context) {
	ctx := c.Request.Context()
	meetings, err := mh.meetingService.GetMeetings(ctx)

	if err != nil {
		mh.log.Error("Failed to get meetings", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	utils.BuildSuccessResponse(c, dto.BuildMeetingToDtoMultiple(meetings))
}

func (mh *MeetingHandler) GetMeetingByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if id == "" {
		mh.log.Error("Missing ID in request")
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	meeting, err := mh.meetingService.GetMeetingByID(ctx, id)
	if err != nil {
		mh.log.Error("Failed to get meeting by ID", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	if meeting == nil {
		utils.BuildErrorResponse(c, http.StatusNotFound, domain_errors.MeetingNotFound)
		return
	}

	utils.BuildSuccessResponse(c, dto.BuildMeetingToDto(*meeting))
}

func (mh *MeetingHandler) DeleteMeetingByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if id == "" {
		mh.log.Error("Missing ID in request")
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	err := mh.meetingService.DeleteMeetingByID(ctx, id)
	if err != nil {
		mh.log.Error("Failed to delete meeting by ID", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	utils.BuildSuccessResponse(c, gin.H{"message": "Meeting deleted successfully"})
}
